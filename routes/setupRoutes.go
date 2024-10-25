package routes

import (
	"forum/handlers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	
	// page handlers
	mux.HandleFunc("/assets/", handlers.AssetsHandler)
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)

	// api handlers
	mux.HandleFunc("/users", handlers.PostUser)
	mux.HandleFunc("/login", handlers.LoginUser)
	mux.HandleFunc("/check_session", handlers.SessionHandler)

	return mux
}
