package handlers

import (
	"calculator-go/models"
	"calculator-go/utils"
	"encoding/json"
	"net/http"
)

var users = make(map[string]string)
var userIdSeq uint = 1

func Register(res http.ResponseWriter, req *http.Request) {
	var user models.User

	// Parse le corps de la requête JSON
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		utils.HandleError(res, req, http.StatusBadRequest, "Invalid data")
		return
	}

	// Vérifie si l'utilisateur existe déjà
	if _, ok := users[user.Username]; ok {
		utils.HandleError(res, req, http.StatusBadRequest, "User already exists")
		return
	}

	// Ajoute l'utilisateur à la liste
	users[user.Username] = user.Password
	user.Id = userIdSeq
	userIdSeq++

	utils.TextResponse(res, req, http.StatusCreated, "User created")
}

func Login(res http.ResponseWriter, req *http.Request) {
	var user models.User

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		utils.HandleError(res, req, http.StatusUnauthorized, "Invalid data")
		return
	}

	if _, ok := users[user.Username]; !ok {
		utils.HandleError(res, req, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := utils.GenerateToken(int(user.Id))
	if err != nil {
		utils.HandleError(res, req, http.StatusBadRequest, "Token generation failed")
	}

	utils.TextResponse(res, req, http.StatusOK, token)
}
