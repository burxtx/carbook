package http

import (
	"context"
	"encoding/json"
	"net/http"

	mylog "github.com/burxtx/car/libs/log"
	"github.com/burxtx/car/posts"
	"github.com/burxtx/car/posts/writing"
	"github.com/go-chi/chi"
)

type writingHandler struct {
	s      writing.Service
	logger mylog.MyLogger
}

func (h *writingHandler) router() chi.Router {
	r := chi.NewRouter()
	r.Route("/writing", func(r chi.Router) {
		r.Post("/new", h.Create)
	})
	return r
}

func (h *writingHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var request struct {
		author  string
		vechile string
		title   string
		excerpt string
		body    string
	}
	p, err := h.s.Create(ctx, request.author, request.vechile, request.title, request.excerpt, request.body)
	if err != nil {
		h.logger.GetLogger(ctx).Errorln(err.Error())
		encodeError(ctx, err, w)
		return
	}
	var response = struct {
		ID posts.PostID `json:"post_id"`
	}{
		ID: p.ID,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.GetLogger(ctx).Errorln(err.Error())
		encodeError(ctx, err, w)
		return
	}
}
