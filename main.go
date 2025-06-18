package main

import (
	"gabtec/go-log-producer/internal/handler"
	"gabtec/go-log-producer/internal/mw"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := ":4000"

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	slog.Info("Server started")

	// ROUTES
	mux := http.NewServeMux()

	// GET /  <--- MUST be last one
	mux.HandleFunc("/", handler.IndexHandler)

	// GET /random
	mux.HandleFunc("/random", handler.RandomHandler)

	// GET /log/:type
	mux.HandleFunc("/log/{type}", handler.LogHandler)

	// for retrocompatibility
	mux.HandleFunc("/demo", handler.RandomHandler)

	// MW
	muxWithMW := mw.ReqLogger(mux)

	slog.Info("Starting server on " + addr)
	http.ListenAndServe(addr, muxWithMW)
}
