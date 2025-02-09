package handlers

import (
	"encoding/json"
	"net/http"

	"go-microservices/internal/models"
	"go-microservices/internal/services"
)

var accountService = services.NewAccountService()

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdAccount := accountService.CreateAccount(account.Name, account.Balance)
	if createdAccount == nil {
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdAccount)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	// Implementation for retrieving an account
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an account
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting an account
}
