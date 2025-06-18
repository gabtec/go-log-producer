package builder

import (
	"encoding/json"
	"gabtec/go-log-producer/internal/model"
	"gabtec/go-log-producer/internal/utils"
	"log/slog"
	"net/http"
)

func BuildSuccessLog(index int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")

	slog.Info("Generating a success log")
	resp := map[string]string{
		"success": "ok",
		"data":    utils.GetSuccessMessage(index),
	}
	slog.Info(resp["data"])
	json.NewEncoder(w).Encode(resp)
}

func BuildErrorLog(code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")

	slog.Info("Generating an error log")
	resp := model.ApiError{
		StatusCode: code,
		StatusText: http.StatusText(code),
		Message:    utils.GetErrorMessage(code),
	}
	slog.Error(resp.Message)
	json.NewEncoder(w).Encode(resp)
}
