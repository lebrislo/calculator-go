package middleware

import (
	"fmt"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Auth Hook !")
		next.ServeHTTP(res, req)
	})
}
