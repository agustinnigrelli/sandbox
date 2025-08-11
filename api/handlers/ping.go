package handlers

import (
	"net/http"

	"github.com/agustinnigrelli/sandbox/utils/web"
)

type Ping interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type PingHandler struct{}

func NewPingHandler() Ping {
	return &PingHandler{}
}

func (h *PingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	web.JsonResponse(w, http.StatusOK, map[string]string{"message": "pong"})
}
