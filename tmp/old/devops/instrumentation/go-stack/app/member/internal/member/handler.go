package member

import (
	"net/http"

	"example.com/member/internal/tracing"
	"github.com/agussyahrilmubarok/gox/pkg/xexception"
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
	ctx, span := tracing.StartSpan(ctx, "handler.FindAll")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.FindAll").Logger()

	members, err := h.service.FindAll(ctx)
	if err != nil {
		span.RecordError(err)
		log.Error().Err(err).Msg("failed to fetch members")
		if httpErr, ok := err.(*xexception.Http); ok {
			return c.JSON(httpErr.Code, map[string]string{"error": httpErr.Message})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch members"})
	}

	log.Info().Int("member_count", len(members)).Msg("successfully retrieved all members")
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
	ctx, span := tracing.StartSpan(ctx, "handler.Find")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.Find").Logger()

	memberCode := c.QueryParam("code")
	if memberCode == "" {
		log.Warn().Msg("missing code parameter")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "missing code parameter"})
	}

	member, err := h.service.FindByCode(ctx, memberCode)
	if err != nil {
		span.RecordError(err)
		log.Error().Err(err).Str("member_code", memberCode).Msg("failed to fetch member")
		if httpErr, ok := err.(*xexception.Http); ok {
			return c.JSON(httpErr.Code, map[string]string{"error": httpErr.Message})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch member"})
	}

	log.Info().Str("member_code", memberCode).Msg("successfully retrieved member")
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
	ctx := c.Request().Context()
	ctx, span := tracing.StartSpan(ctx, "handler.InitDummy")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.InitDummy").Logger()

	dummies := []*Member{
		{ID: uuid.New().String(), Code: "MEMBER-1000", Name: "John Doe", Email: "johndoe@mail.com"},
		{ID: uuid.New().String(), Code: "MEMBER-1001", Name: "Jane Smith", Email: "janesmith@mail.com"},
	}

	for _, m := range dummies {
		if _, err := h.service.Save(ctx, m); err != nil {
			span.RecordError(err)
			log.Error().Err(err).Str("member_id", m.ID).Msg("failed to insert dummy member")
			if httpErr, ok := err.(*xexception.Http); ok {
				return c.JSON(httpErr.Code, map[string]string{
					"error": httpErr.Message,
				})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to insert dummy member"})
		}
	}

	log.Info().Msg("dummy data initialized")
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
	ctx := c.Request().Context()
	ctx, span := tracing.StartSpan(ctx, "handler.CleanDummy")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.CleanDummy").Logger()

	members, err := h.service.FindAll(ctx)
	if err != nil {
		span.RecordError(err)
		log.Error().Err(err).Msg("failed to fetch members")
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch members"})
	}

	for _, m := range members {
		if err := h.service.DeleteByID(ctx, m.ID); err != nil {
			span.RecordError(err)
			log.Error().Err(err).Str("member_id", m.ID).Str("member_code", m.Code).Msg("failed to delete dummy member")
			if httpErr, ok := err.(*xexception.Http); ok {
				return c.JSON(httpErr.Code, map[string]string{
					"error": httpErr.Message,
				})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to delete dummy member"})
		}
	}

	log.Info().Msg("dummy data cleaned")
	return c.JSON(http.StatusOK, echo.Map{"message": "dummy data cleaned"})
}
