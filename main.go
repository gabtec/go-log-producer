package main

import (
	"encoding/json"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
)

type ApiError struct {
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
}

func main() {
	addr := ":4000"

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	slog.Info("Server started")

	http.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
		randomIndex := rand.Intn(6)
		slog.Info("Random integer generated ok")

		code := getHttpCode(randomIndex)
		w.WriteHeader(code)
		w.Header().Add("Content-Type", "application/json")

		if code == http.StatusOK {
			slog.Info("Success code was selected")
			resp := map[string]string{
				"success": "ok",
				"data":    getSuccessMessage(randomIndex),
			}
			slog.Info(resp["data"])
			json.NewEncoder(w).Encode(resp)

		} else {
			slog.Warn("Error code was selected")
			resp := ApiError{
				StatusCode: code,
				StatusText: http.StatusText(code),
				Message:    getErrorMessage(code),
			}
			slog.Error(resp.Message)
			json.NewEncoder(w).Encode(resp)
		}

		// w.Write([]byte("Demo OK"))
	})

	slog.Info("Starting server on " + addr)
	http.ListenAndServe(addr, nil)
}

func getHttpCode(idx int) int {
	possibleResults := []int{200, 400, 401, 403, 404, 500}
	return possibleResults[idx]
}

func getErrorMessage(code int) string {
	messages := map[int]string{
		400: "Account not found",
		401: "You are not authorized to create users",
		403: "You have no permissions to delete accounts",
		404: "Missing required field: 'X'",
		500: "Unexpected server behavior",
	}

	return messages[code]
}

func getSuccessMessage(idx int) string {
	messages := []string{
		"Object created successfully",
		"Object updated successfully",
		"Object deleted successfully",
		"User X loggedIn successfully",
		"Cronjob executed successfully",
	}

	return messages[idx]
}
