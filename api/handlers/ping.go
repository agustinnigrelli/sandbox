package handlers

import (
	"net/http"

	"github.com/agustinnigrelli/sandbox/utils/web"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	web.JsonResponse(w, http.StatusOK, map[string]string{"message": "pong"})
}
