package main

import (
	"fmt"
	"net/http"

	"group-tracker/handlers"
)

func main() {
	// Routes
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/moreinfo", handlers.Artist)

	// Static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
