package main

import (
	"log"
	"net/http"
	"os"

	"calculator-go/handlers"
	"calculator-go/middleware"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {
	// Configure logger
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:  true,
		PadLevelText: false,
	})

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register routes with logging middleware
	mux.Handle("/add", middleware.LoggingMiddleware(http.HandlerFunc(handlers.Add)))
	mux.Handle("/sub", middleware.LoggingMiddleware(http.HandlerFunc(handlers.Sub)))
	mux.Handle("/mul", middleware.LoggingMiddleware(http.HandlerFunc(handlers.Mul)))
	mux.Handle("/div", middleware.LoggingMiddleware(http.HandlerFunc(handlers.Div)))

	// Apply auth middleware globally
	authMux := middleware.Auth(mux)
	// Apply rateLimiter middleware globally
	rateLimitedMux := middleware.RateLimiter(authMux)

	log.Fatal(http.ListenAndServe(":8080", rateLimitedMux))

	logger.Info("Server is starting on port 8080...")
}
