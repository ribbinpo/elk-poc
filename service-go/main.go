package main

import (
	"io"
	"log"
	"log/slog"
	"os"

	logs "github.com/ribbinpo/elk-poc/service-go/pkg/logs"
)

func main() {
	// console
	consoleWriter := os.Stdout
	// create file
	file, err := os.OpenFile("logs/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	httpWriter := &logs.HttpLogWriter{
		URL: "http://localhost:8080",
	}

	multiWriter := io.MultiWriter(consoleWriter, file, httpWriter)
	logger := slog.New(slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{}))

	logger.Info("hello", "count", 4) // slog.Group("properties",
	// 	slog.Int("width", 640),
	// 	slog.Int("height", 480),
	// 	slog.String("format", "jpeg"),
	// )

}
