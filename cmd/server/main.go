package main

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/FilipSolich/go-template/internal/log"
	"github.com/FilipSolich/go-template/internal/version"
)

func main() {
	logger, zLogger := log.NewDevelopment()
	defer zLogger.Sync()

	slog.SetDefault(logger)
	info := version.Info()
	slog.Info("Starting server",
		slog.String("version", info.Version),
		slog.String("commit", info.Commit),
		slog.String("goVersion", info.GoVersion),
	)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	})
	if err := http.ListenAndServe(":8000", nil); err != nil {
		slog.Error("Failed to run server", log.Err(err))
	}
}
