package utils

import "net/http"

func AddCookie(w http.ResponseWriter, name, value string) {
	http.SetCookie(w, &http.Cookie{
		Name:   name,
		Value:  value,
		Path:   "/",
		MaxAge: 60*60*24,
	})
}

func DeleteCookie(w http.ResponseWriter, session_name string) {
	http.SetCookie(w, &http.Cookie{
		Name:   session_name,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
