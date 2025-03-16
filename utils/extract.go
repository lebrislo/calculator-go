package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

// ExtractNumbers retrieves "a" and "b" parameters from the request
func ExtractNumbers(req *http.Request) (int, int, error) {
	queryMap := req.URL.Query()

	if !queryMap.Has("a") || !queryMap.Has("b") {
		return 0, 0, fmt.Errorf("missing request parameters")
	}

	aStr := queryMap.Get("a")
	bStr := queryMap.Get("b")

	aNum, aErr := strconv.Atoi(aStr)
	bNum, bErr := strconv.Atoi(bStr)

	if aErr != nil || bErr != nil {
		return 0, 0, fmt.Errorf("invalid input. Please provide integers for 'a' and 'b'")
	}

	return aNum, bNum, nil
}
