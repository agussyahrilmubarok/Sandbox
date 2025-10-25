package member

import (
	"context"
	"net/http"

	"github.com/google/uuid"
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

// FindAll godoc
// @Summary Get all members
// @Description Retrieves all member data
// @Tags Members
// @Produce json
// @Success 200 {array} Member
// @Failure 500 {object} map[string]string
// @Router /members [get]
func (h *Handler) FindAll(c echo.Context) error {
	ctx := c.Request().Context()
	log := zerolog.Ctx(ctx)

	members, err := h.service.FindAll(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch members")
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch members"})
	}

	log.Info().Int("member_count", len(members)).Msg("Successfully retrieved all members")
	return c.JSON(http.StatusOK, members)
}

// Find godoc
// @Summary Get member by query param
// @Description Retrieves member data by query param
// @Tags Members
// @Produce json
// @Param code query string true "Member Code"
// @Success 200 {object} Member
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /members/find [get]
func (h *Handler) Find(c echo.Context) error {
	ctx := c.Request().Context()
	log := zerolog.Ctx(ctx)

	memberCode := c.QueryParam("code")

	if memberCode == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing code parameter"})
	}

	member, err := h.service.FindByCode(ctx, memberCode)
	if err != nil {
		log.Error().Err(err).Str("member_code", memberCode).Msg("Failed to fetch member")
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch member"})
	}

	log.Info().Str("member_code", memberCode).Msg("Successfully retrieved member")
	return c.JSON(http.StatusOK, member)
}

// InitDummy godoc
// @Summary Initialize dummy member data
// @Description Create dummy data for testing
// @Tags Dummy
// @Produce json
// @Success 200 {array} Member
// @Failure 500 {object} map[string]string
// @Router /members/init-dummy [post]
func (h *Handler) InitDummy(c echo.Context) error {
	ctx := context.Background()
	log := zerolog.Ctx(ctx)

	dummies := []*Member{
		{ID: uuid.New().String(), Code: "MC-1XX", Name: "John Doe", Email: "johndoe@mail.com"},
		{ID: uuid.New().String(), Code: "MC-2XX", Name: "Jane Smith", Email: "janesmith@mail.com"},
	}

	for _, m := range dummies {
		if _, err := h.service.Save(ctx, m); err != nil {
			log.Error().Err(err).Str("member_id", m.ID).Msg("Failed to insert dummy member")
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to insert dummy member"})
		}
	}

	log.Info().Msg("Dummy data initialized")
	return c.JSON(http.StatusOK, dummies)
}

// CleanDummy godoc
// @Summary Remove all dummy members
// @Description Delete all dummy member data
// @Tags Dummy
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /members/clean-dummy [delete]
func (h *Handler) CleanDummy(c echo.Context) error {
	ctx := context.Background()
	log := zerolog.Ctx(ctx)

	members, err := h.service.FindAll(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch members")
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch members"})
	}

	for _, m := range members {
		if err := h.service.DeleteByID(ctx, m.ID); err != nil {
			log.Error().Err(err).Str("member_id", m.ID).Msg("Failed to delete dummy member")
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete dummy member"})
		}
	}

	log.Info().Msg("Dummy data cleaned")
	return c.JSON(http.StatusOK, echo.Map{"message": "Dummy data cleaned"})
}
