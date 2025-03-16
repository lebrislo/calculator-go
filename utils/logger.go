package utils

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// LogRequest logs request details
func LogRequest(r *http.Request) {
	requestId := r.Header.Get("requestId")
	logger.WithFields(logrus.Fields{
		"method":    r.Method,
		"url":       r.URL.Path,
		"params":    r.URL.Query(),
		"host":      r.Host,
		"requestId": requestId,
	}).Info("Request received")
}

// LogResponse logs response details
func LogResponse(r *http.Request, status int, message string, result int) {
	requestId := r.Header.Get("requestId")
	entry := logger.WithFields(logrus.Fields{
		"method":    r.Method,
		"url":       r.URL.Path,
		"params":    r.URL.Query(),
		"host":      r.Host,
		"requestId": requestId,
		"status":    status,
	})

	if status >= 200 && status < 300 {
		entry = entry.WithField("result", result)
		entry.Info(message)
	} else {
		entry.Error(message)
	}
}
