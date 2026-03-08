package booking

import (
	"net/http"

	"example.com/booking/pkg/exception"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}

// Booking godoc
// @Summary Book a course seat
// @Description Book a seat for a course by providing booking details
// @Tags Booking
// @Accept json
// @Produce json
// @Param booking_request body BookingRequest true "Booking Request"
// @Success 200 {object} Booking "Successfully booked the course"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/booking/course [post]
func (h *Handler) Booking(c echo.Context) error {
	ctx := c.Request().Context()
	log := zerolog.Ctx(ctx)

	var payload BookingRequest
	if err := c.Bind(&payload); err != nil {
		log.Error().Err(err).Msg("Failed to decode request body")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}

	booking, err := h.service.Booking(ctx, payload)
	if err != nil {
		log.Error().
			Err(err).
			Str("member_code", payload.MemberCode).
			Str("course_code", payload.CourseCode).
			Msg("Failed to book course")

		if httpErr, ok := err.(*exception.Http); ok {
			return c.JSON(httpErr.Code, map[string]string{
				"error": httpErr.Message,
			})
		}

		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to book course"})
	}

	log.Info().
		Str("member_code", payload.MemberCode).
		Str("course_code", payload.CourseCode).
		Msg("Course booked successfully")

	return c.JSON(http.StatusOK, booking)
}

// BookingV2 godoc
// @Summary Book a course seat (v2)
// @Description Book a seat for a course using goroutines (parallel processing)
// @Tags Booking
// @Accept json
// @Produce json
// @Param booking_request body BookingRequest true "Booking Request"
// @Success 200 {object} Booking "Successfully booked the course"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v2/booking/course [post]
func (h *Handler) BookingV2(c echo.Context) error {
	ctx := c.Request().Context()
	log := zerolog.Ctx(ctx)

	var payload BookingRequest
	if err := c.Bind(&payload); err != nil {
		log.Error().Err(err).Msg("Failed to decode request body")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}

	booking, err := h.service.Booking(ctx, payload)
	if err != nil {
		log.Error().
			Err(err).
			Str("member_code", payload.MemberCode).
			Str("course_code", payload.CourseCode).
			Msg("Failed to book course")

		if httpErr, ok := err.(*exception.Http); ok {
			return c.JSON(httpErr.Code, map[string]string{
				"error": httpErr.Message,
			})
		}

		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to book course"})
	}

	log.Info().
		Str("member_code", payload.MemberCode).
		Str("course_code", payload.CourseCode).
		Msg("Course booked successfully")

	return c.JSON(http.StatusOK, booking)
}
