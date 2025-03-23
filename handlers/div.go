package handlers

import (
	"calculator-go/utils"
	"fmt"
	"net/http"
)

func Div(res http.ResponseWriter, req *http.Request) {
	a, b, err := utils.ExtractNumbers(req)
	if err != nil {
		utils.HandleError(res, req, http.StatusBadRequest, err.Error())
		return
	}

	if b == 0 {
		utils.HandleError(res, req, http.StatusBadRequest, "Division by 0 forbidden")
		return
	}

	value := a / b
	utils.SendResult(res, value)
	utils.LogResponse(req, http.StatusOK, "Division performed", fmt.Sprint(value))
}
