package member

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Handler struct {
	service IService
	log     zerolog.Logger
}

func NewHandler(service IService, log zerolog.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log.With().Str("component", "member_handler").Logger(),
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
func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	members, err := h.service.FindAll(ctx)
	if err != nil {
		h.log.Error().Err(err).Msg("Failed to fetch members")
		http.Error(w, "Failed to fetch members", http.StatusInternalServerError)
		return
	}

	h.log.Info().Int("member_count", len(members)).Msg("Successfully retrieved all members")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

// FindByCode godoc
// @Summary Get member by Code
// @Description Retrieves member data by code
// @Tags Members
// @Produce json
// @Param code query string true "Member Code"
// @Success 200 {object} Member
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /members/find [get]
func (h *Handler) FindByCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	memberCode := r.URL.Query().Get("code")
	if memberCode == "" {
		http.Error(w, "Missing code parameter", http.StatusBadRequest)
		return
	}

	member, err := h.service.FindByCode(ctx, memberCode)
	if err != nil {
		h.log.Error().Err(err).Str("member_code", memberCode).Msg("Failed to fetch member")
		http.Error(w, "Failed to fetch member", http.StatusInternalServerError)
		return
	}

	if member == nil {
		http.Error(w, "Member not found", http.StatusNotFound)
		return
	}

	h.log.Info().Str("member_code", memberCode).Msg("Successfully retrieved member")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(member)
}

// InitDummy godoc
// @Summary Initialize dummy member data
// @Description Create dummy data for testing
// @Tags Dummy
// @Produce json
// @Success 200 {array} Member
// @Router /members/init-dummy [post]
func (h *Handler) InitDummy(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	dummies := []*Member{
		{ID: uuid.New().String(), Code: "MC-1XX", Name: "John Doe", Email: "johndoe@mail.com"},
		{ID: uuid.New().String(), Code: "MC-2XX", Name: "Jane Smith", Email: "janesmith@mail.com"},
	}

	for _, m := range dummies {
		if _, err := h.service.Save(ctx, m); err != nil {
			h.log.Error().Err(err).Str("member_id", m.ID).Msg("Failed to insert dummy member")
		}
	}

	h.log.Info().Msg("Dummy data initialized")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dummies)
}

// CleanDummy godoc
// @Summary Remove all dummy members
// @Description Delete all dummy member data
// @Tags Dummy
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /members/clean-dummy [delete]
func (h *Handler) CleanDummy(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	members, err := h.service.FindAll(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch members", http.StatusInternalServerError)
		return
	}

	for _, m := range members {
		if err := h.service.DeleteByID(ctx, m.ID); err != nil {
			h.log.Error().Err(err).Str("member_id", m.ID).Msg("Failed to delete dummy member")
		}
	}

	h.log.Info().Msg("Dummy data cleaned")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Dummy data cleaned"})
}
