package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	endpoints := map[string]string{
		"/":          "this page",
		"/log/:type": "returns a success (type=ok), or error (type=error), response",
		"/random":    "returns a random response",
		"/demo":      "same as /random (for back compatibility)",
	}

	resp := map[string]map[string]string{
		"endpointsList": endpoints,
	}
	slog.Info("Index endpoint visited")
	json.NewEncoder(w).Encode(resp)
}
