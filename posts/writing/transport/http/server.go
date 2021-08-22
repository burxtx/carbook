package http

import (
	"context"
	"encoding/json"
	"net/http"

	mylog "github.com/burxtx/car/libs/log"
	"github.com/burxtx/car/posts/writing"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type HTTPServer struct {
	writingSvc writing.Service

	logger mylog.MyLogger

	router chi.Router
}

func NewHTTPServer(writingSvc writing.Service, logger mylog.MyLogger) *HTTPServer {
	s := &HTTPServer{}
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Route("/post", func(r chi.Router) {
		h := writingHandler{s.writingSvc, logger}
		r.Mount("/v1", h.router())
	})

	r.Method("GET", "/metrics", promhttp.Handler())
	s.router = r
	return s
}

func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
