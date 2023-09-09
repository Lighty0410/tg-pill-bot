package server

import (
	"net/http"
)

type Service interface{}

type HttpHandler[S Service] struct {
	Ctrl S
}

func (h *HttpHandler[S]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/get_pill_time":
		err := h.GetPillTime(w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func NewHttpServer() *http.Server {
	handler := HttpHandler[Service]{
		Ctrl: nil,
	}
	httpServer := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	return &httpServer
}
