package routers

import (
	"net/http"

	"forum/handlers"
	"forum/handlers/api"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /assets/", handlers.AssetsHandler)

	mux.HandleFunc("GET /", handlers.HomeHandler)
	mux.HandleFunc("GET /login", handlers.LoginHandler)
	mux.HandleFunc("GET /register", handlers.RegisterHandler)
	mux.HandleFunc("GET /api/posts", handlers.LoadData)

	mux.HandleFunc("GET /createpost", handlers.CreatePostHandler)

	mux.HandleFunc("POST /api/login", api.LoginUser)
	mux.HandleFunc("POST /api/users", api.PostUser)
	
	// create posts
	mux.HandleFunc("POST /newposts", api.CreatePostsHandler)
	// create comments
	mux.HandleFunc("POST /newcomment", api.CreateCommentsHandler)
}
