package main

import (
	"log"
	"net/http"
	"pbkk-fp-dd-registration-sites/config"
)

func main() {
	config.ConnectDB()
	
	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}