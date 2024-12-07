package homecontroller

import (
	"net/http"
	"text/template"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/login.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}