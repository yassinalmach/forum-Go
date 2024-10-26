package controllers

import (
	"net/http"

	"forum/models"
)

func CreateUser(user models.User) (int, error) {
	return 0, nil
}

func LoginUser(user models.User) (models.User, error) {
	var f models.User
	return f, nil

}

func StoreSession(id string, user models.User) error {
	return nil
}

func GetSession(r *http.Request) (models.User, error) {
	var f models.User
	return f, nil
}
