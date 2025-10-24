package booking

import "github.com/rs/zerolog"

type Handler struct {
	service IService
	log     zerolog.Logger
}

func NewHandler(
	service IService,
	log zerolog.Logger,
) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) Booking() {

}
