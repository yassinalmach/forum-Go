package main

import (
	"fmt"
	"forum/database"
	"forum/routes"
	"log"
	"net/http"
)

func main() {
	db, err := database.SetupDb()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	mux := routes.SetupRoutes()
	
	fmt.Println("Server running on: http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
	
}
