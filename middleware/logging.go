package middleware

import (
	"calculator-go/utils"
	"net/http"
	"strconv"
)

var nextID = make(chan string)

func init() {
	go requestIdCounter()
}

func requestIdCounter() {
	for i := 0; ; i++ {
		nextID <- strconv.Itoa(i)
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("requestId", <-nextID)
		utils.LogRequest(r)
		next.ServeHTTP(w, r)
	})
}
