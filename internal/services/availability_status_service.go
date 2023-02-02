package services

import (
	"net/http"

	"github.com/RHEnVision/provisioning-backend/internal/background"
	"github.com/RHEnVision/provisioning-backend/internal/kafka"
	"github.com/RHEnVision/provisioning-backend/internal/payloads"
	"github.com/go-chi/render"
)

func AvailabilityStatus(w http.ResponseWriter, r *http.Request) {
	payload := &payloads.AvailabilityStatusRequest{}
	if err := render.Bind(r, payload); err != nil {
		renderError(w, r, payloads.NewInvalidRequestError(r.Context(), "availability status", err))
		return
	}

	asm := kafka.AvailabilityStatusMessage{SourceID: payload.SourceID}
	background.EnqueueAvailabilityStatusRequest(&asm)
	writeOk(w, r)
}
