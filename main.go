package main

import (
	"log"
	"net/http"
	"pbkk-fp-dd-registration-sites/config"
	"pbkk-fp-dd-registration-sites/controllers/applicationcontroller"
	"pbkk-fp-dd-registration-sites/controllers/homecontroller"
)

func main() {
	config.ConnectDB()

	// Serve static files (CSS, JS, etc.)
	fs := http.FileServer(http.Dir("views/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.HandleFunc("/", homecontroller.Welcome)
	http.HandleFunc("/login", homecontroller.Login)
	http.HandleFunc("/register", applicationcontroller.Register)
	
	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}