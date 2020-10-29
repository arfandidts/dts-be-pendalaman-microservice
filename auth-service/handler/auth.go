package handler

import (
	"net/http"

	"github.com/arfandidts/dts-be-pendalaman-microservice/utils"
	"gorm.io/gorm"
)

type AuthDB struct {
	DB *gorm.DB
}

func (db *AuthDB) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (db *AuthDB) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func ValidateAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		utils.WrapAPIError(w, r, "Invalid auth", http.StatusForbidden)
		return
	}

	if authToken != "asdfghjk" {
		utils.WrapAPIError(w, r, "Invalid auth", http.StatusForbidden)
		return
	}

	utils.WrapAPISuccess(w, r, "success", 200)
}
