package middleware

import (
	"calculator-go/utils"
	"net/http"
	"strings"
)

func Authentification(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		tokenString := req.Header.Get("Authorization")
		if tokenString == "" {
			utils.HandleError(res, req, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}

		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.HandleError(res, req, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}

		tokenString = tokenParts[1]

		_, err := utils.VerifyToken(tokenString)
		if err != nil {
			utils.HandleError(res, req, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}

		next.ServeHTTP(res, req)
	})
}
