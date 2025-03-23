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

	// Create a new ServeMux for public routes
	publicMux := http.NewServeMux()
	publicMux.Handle("/login", middleware.LoggingMiddleware(http.HandlerFunc(handlers.Login)))
	publicMux.Handle("/register", middleware.LoggingMiddleware(http.HandlerFunc(handlers.Register)))

	// Create a new ServeMux for protected routes
	protectedMux := http.NewServeMux()
	protectedMux.Handle("/add", middleware.LoggingMiddleware(http.HandlerFunc(handlers.Add)))
	protectedMux.Handle("/sub", middleware.LoggingMiddleware(http.HandlerFunc(handlers.Sub)))
	protectedMux.Handle("/mul", middleware.LoggingMiddleware(http.HandlerFunc(handlers.Mul)))
	protectedMux.Handle("/div", middleware.LoggingMiddleware(http.HandlerFunc(handlers.Div)))

	// Apply auth middleware to protected routes
	authProtectedMux := middleware.Authentification(protectedMux)

	// Combine public and protected routes into a single mux
	mainMux := http.NewServeMux()
	mainMux.Handle("/login", publicMux)
	mainMux.Handle("/register", publicMux)
	mainMux.Handle("/add", authProtectedMux)
	mainMux.Handle("/sub", authProtectedMux)
	mainMux.Handle("/mul", authProtectedMux)
	mainMux.Handle("/div", authProtectedMux)

	// Apply rateLimiter middleware globally
	rateLimitedMux := middleware.RateLimiter(mainMux)

	log.Fatal(http.ListenAndServe(":8080", rateLimitedMux))

	logger.Info("Server is starting on port 8080...")
}
