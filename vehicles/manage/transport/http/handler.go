package http

import (
	"context"
	"encoding/json"
	"net/http"

	mylog "github.com/burxtx/car/libs/log"
	"github.com/burxtx/car/vehicles"
	"github.com/burxtx/car/vehicles/manage"
	"github.com/go-chi/chi"
)

type Handler struct {
	s      manage.Service
	logger mylog.MyLogger
}

func (h *Handler) router() chi.Router {
	r := chi.NewRouter()
	r.Route("/manage", func(r chi.Router) {
		r.Post("/new", h.Create)
	})
	return r
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var request struct {
		name      string
		brand     string
		modelName string
		year      string
	}
	v, err := h.s.Create(ctx, request.name, request.brand, request.modelName, request.year)
	if err != nil {
		h.logger.GetLogger(ctx).Errorln(err.Error())
		encodeError(ctx, err, w)
		return
	}
	var response = struct {
		ID vehicles.VehicleID `json:"vehicle_id"`
	}{
		ID: v.ID,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.GetLogger(ctx).Errorln(err.Error())
		encodeError(ctx, err, w)
		return
	}
}
