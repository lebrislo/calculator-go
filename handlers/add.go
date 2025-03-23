package handlers

import (
	"calculator-go/utils"
	"fmt"
	"net/http"
)

func Add(res http.ResponseWriter, req *http.Request) {
	a, b, err := utils.ExtractNumbers(req)
	if err != nil {
		utils.HandleError(res, req, http.StatusBadRequest, err.Error())
		return
	}

	value := a + b
	utils.SendResult(res, value)
	utils.LogResponse(req, http.StatusOK, "Addition performed", fmt.Sprint(value))
}
