package main

import (
	"log"
	"net/http"
	"pbkk-fp-dd-registration-sites/config"
	"pbkk-fp-dd-registration-sites/controllers/admindashboard"
	"pbkk-fp-dd-registration-sites/controllers/adminuniversity"
	"pbkk-fp-dd-registration-sites/controllers/applicationcontroller"
	"pbkk-fp-dd-registration-sites/controllers/homecontroller"
	"pbkk-fp-dd-registration-sites/controllers/userdashboard"
)

func main() {
	config.ConnectDB()

	// Serve static files (CSS, JS, etc.)
	fs := http.FileServer(http.Dir("views/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	// Homepage
	http.HandleFunc("/", homecontroller.Welcome)
	http.HandleFunc("/login", homecontroller.Login)
	
	// Application
	http.HandleFunc("/register", applicationcontroller.Register)

	// User
	http.HandleFunc("/user-dash", userdashboard.Dashboard)

	// Admin
	http.HandleFunc("/admin-dash", admindashboard.Dashboard)
	http.HandleFunc("/admin-uni", adminuniversity.Dashboard)


	


	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}