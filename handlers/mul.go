package handlers

import (
	"calculator-go/utils"
	"net/http"
)

func Mul(res http.ResponseWriter, req *http.Request) {
	a, b, err := utils.ExtractNumbers(req)
	if err != nil {
		utils.HandleError(res, req, http.StatusBadRequest, err.Error())
		return
	}

	value := a * b
	utils.SendResult(res, value)
	utils.LogResponse(req, http.StatusOK, "Multiplication performed", value)
}
