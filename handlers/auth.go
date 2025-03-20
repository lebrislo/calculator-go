package handlers

import (
	"calculator-go/models"
	"calculator-go/utils"
	"encoding/json"
	"net/http"
)

func Register(res http.ResponseWriter, req *http.Request) {
	var user models.User

	// Parse le corps de la requête JSON
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		utils.HandleError(res, req, http.StatusUnauthorized, "Invalid data")
		return
	}

	user.Id = 1
	utils.LogResponse(req, http.StatusCreated, "User registered successfully", 0)
}

func Login(res http.ResponseWriter, req *http.Request) {

}
