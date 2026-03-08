package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/main-service/internal/domain"
	"example.com/main-service/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type flightHandler struct {
	uc        usecase.IFlightUseCase
	log       *zap.Logger
	validator *validator.Validate
}

func NewFlightHandler(uc usecase.IFlightUseCase, log *zap.Logger) *flightHandler {
	return &flightHandler{
		uc:        uc,
		log:       log,
		validator: validator.New(),
	}
}

func (h *flightHandler) SearchFlights(c *fiber.Ctx) error {
	var body domain.CreateSearchBody
	if err := c.BodyParser(&body); err != nil {
		h.log.Error("invalid request body", zap.Error(err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	if err := h.validator.Struct(body); err != nil {
		h.log.Warn("validation failed", zap.Error(err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Validation error",
			"errors":  err.Error(),
		})
	}

	searchID, err := h.uc.SearchFlights(c.Context(), body)
	if err != nil {
		h.log.Error("Failed to search flights", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to initiate search",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Search request submitted",
		"data": fiber.Map{
			"search_id": searchID,
			"status":    "processing",
		},
	})
}

func (h *flightHandler) StreamFlightResults(c *fiber.Ctx) error {
	searchID := c.Params("search_id")
	if searchID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "missing search_id param",
		})
	}

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	resultsChan := h.uc.StreamResults(c.Context(), searchID)
	fctx := c.Context() // simpan fasthttp.RequestCtx

	fctx.SetBodyStreamWriter(func(w *bufio.Writer) {
		for {
			select {
			case <-fctx.Done():
				h.log.Info("SSE disconnected",
					zap.String("search_id", searchID),
				)
				return
			case res, ok := <-resultsChan:
				if !ok {
					h.log.Info("SSE stream ended",
						zap.String("search_id", searchID),
					)
					return
				}
				h.log.Info("SSE send",
					zap.String("search_id", res.SearchID),
					zap.String("status", res.Status),
					zap.Int("total_results", len(res.Results)),
				)
				data, err := json.Marshal(res)
				if err != nil {
					h.log.Error("failed to marshal result", zap.Error(err))
					continue
				}
				fmt.Fprintf(w, "data: %s\n", data)
				w.Flush()
				if res.Status == "completed" {
					summary := struct {
						SearchID     string `json:"search_id"`
						Status       string `json:"status"`
						TotalResults int    `json:"total_results"`
					}{
						SearchID:     res.SearchID,
						Status:       res.Status,
						TotalResults: len(res.Results),
					}

					data, _ := json.Marshal(summary)
					fmt.Fprintf(w, "data: %s\n\n", data)
					w.Flush()

					h.log.Info("SSE completed", zap.String("search_id", searchID))
					return
				}
			}
		}
	})

	return nil
}
