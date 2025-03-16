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

	logger.Info("Server is starting on port 8080...")

	// Register routes with logging middleware
	http.Handle("/add", middleware.LoggingMiddleware(http.HandlerFunc(handlers.Add)))
	http.Handle("/sub", middleware.LoggingMiddleware(http.HandlerFunc(handlers.Sub)))
	http.Handle("/mul", middleware.LoggingMiddleware(http.HandlerFunc(handlers.Mul)))
	http.Handle("/div", middleware.LoggingMiddleware(http.HandlerFunc(handlers.Div)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
