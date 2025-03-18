package middleware

import (
	"calculator-go/utils"
	"net/http"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 3)

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if !limiter.Allow() {
			utils.HandleError(res, req, http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))
			return
		}

		next.ServeHTTP(res, req)
	})
}

/*
Another implementation using channels

var rateLimit = 1 * time.Second
var tokens = make(chan struct{}, 3)

func init() {
	go func() {
		ticker := time.NewTicker(rateLimit)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case tokens <- struct{}{}:
			default:
			}
		}
	}()
}

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		select {
		case <-tokens:
			next.ServeHTTP(res, req)
		default:
			HandleError(res, req, http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))
		}
	})
}
*/
