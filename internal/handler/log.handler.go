package handler

import (
	"gabtec/go-log-producer/internal/builder"
	"gabtec/go-log-producer/internal/utils"
	"log/slog"
	"math/rand"
	"net/http"
)

func LogHandler(w http.ResponseWriter, r *http.Request) {
	selectedLogType := r.PathValue("type")
	slog.Info("Selected log type: " + selectedLogType)
	slog.Warn("--debug")
	slog.Warn(selectedLogType)
	randomIndex := rand.Intn(6)

	switch selectedLogType {
	case "error":
		code := utils.GetHttpErrorCode(randomIndex)
		builder.BuildErrorLog(code, w)
	default:
		builder.BuildSuccessLog(randomIndex, w)
	}
}
