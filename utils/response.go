package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CalculationResponse represents the JSON structure
type CalculationResponse struct {
	Result int `json:"result"`
}

// SendResult sends a JSON response to the client
func SendResult(res http.ResponseWriter, value int) {
	res.Header().Set("Content-Type", "application/json")

	resp := CalculationResponse{Result: value}
	j, _ := json.Marshal(&resp)

	fmt.Fprintln(res, string(j))
}

// HandleError handles error responses
func HandleError(res http.ResponseWriter, req *http.Request, status int, message string) {
	http.Error(res, message, status)
	LogResponse(req, status, message, 0)
}
