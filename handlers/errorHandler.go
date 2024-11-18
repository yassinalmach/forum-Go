package handlers

import (
	"html/template"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	tem, err := template.ParseFiles("./web/templates/error.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	var DetailsError struct {
		Title  string
		Status int
		Method string
		Path   string
	}

	switch status {
	case http.StatusNotFound:
		DetailsError.Title = "Error 404 (Not Found)!!"
	case http.StatusMethodNotAllowed:
		DetailsError.Title = "Error 405 (Method Not Allowed)!!"
	case http.StatusInternalServerError:
		DetailsError.Title = "Error 500 (Internal Server Error)!!"
	case http.StatusCreated:
		DetailsError.Title = "Error 201 (Status Created)!!"
	case http.StatusBadRequest:
		DetailsError.Title = "Error 400 (Status Bad Request)"
	case http.StatusBadGateway:
		DetailsError.Title = "Error 502 (Status Bad Gateway)"
	}

	w.WriteHeader(status)
	DetailsError.Status = status
	DetailsError.Method = r.Method
	DetailsError.Path = r.URL.Path
	if err := tem.Execute(w, DetailsError); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}
