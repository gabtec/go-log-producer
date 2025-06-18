package handler

import (
	"gabtec/go-log-producer/internal/builder"
	"gabtec/go-log-producer/internal/utils"
	"log/slog"
	"math/rand"
	"net/http"
)

func RandomHandler(w http.ResponseWriter, r *http.Request) {
	randomIndex := rand.Intn(6)
	slog.Info("Random integer generated ok")

	code := utils.GetHttpCode(randomIndex)
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")

	if code == http.StatusOK {
		builder.BuildSuccessLog(randomIndex, w)
	} else {
		builder.BuildErrorLog(code, w)
	}

	// w.Write([]byte("Demo OK"))
}
