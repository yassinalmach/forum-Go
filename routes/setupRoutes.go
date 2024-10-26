package routes

import (
	"forum/handlers"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	// asset handler
	mux.HandleFunc("/assets/", handlers.AssetsHandler)

	// page handlers
	mux.HandleFunc("GET /", handlers.HomeHandler)
	mux.HandleFunc("GET /register", handlers.RegisterHandler)
	mux.HandleFunc("GET /login", handlers.LoginHandler)

	// api handlers
	mux.HandleFunc("POST /api/register", handlers.LoginUser)
	mux.HandleFunc("POST /api/login", handlers.LoginUser)
	mux.HandleFunc("POST /api/post", handlers.PostUser)
	mux.HandleFunc("POST /api/comment", handlers.LoginUser)
	mux.HandleFunc("api//check_session", handlers.SessionHandler)
}

