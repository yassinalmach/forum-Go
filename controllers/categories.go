package controllers

import (
	"errors"

	"forum/models"
	"forum/utils"
)

func CreateCategories(categories models.Categories) (bool, error) {
	var count int
	err := utils.DataBase.QueryRow("SELECT COUNT(*) FROM categories WHERE categori = ? nameCategori", categories.Categori).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return false, errors.New("categori already exist")
	}
	stmt, err := utils.DataBase.Prepare("INSERT INTO categories (categori,post_id) VALUES(?,?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(categories.Categori, categories.PostId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func DisplayCategories(categories models.Categories) (string, error) {
	var resalt string
	err := utils.DataBase.QueryRow("SELECT categori FROM categories").Scan(&resalt)
	if err != nil {
		return "", err
	}
	return resalt, nil
}
