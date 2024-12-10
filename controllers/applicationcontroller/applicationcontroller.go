package applicationcontroller

import (
	"html/template"
	"net/http"
	"pbkk-fp-dd-registration-sites/models/admindashboardmodel"
	"pbkk-fp-dd-registration-sites/models/applicationmodel"
)

func Register(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/register.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

func UserData (w http.ResponseWriter, r *http.Request) {
	users := applicationmodel.GetUser()
	data := map[string]any {
		"users": users,
	}

	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func UniDegreeData (w http.ResponseWriter, r *http.Request) {
	unidegrees, err := admindashboardmodel.GetUniDegree()
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	
	data := map[string]any {
		"unidegrees": unidegrees,
	}
	
	temp, err := template.ParseFiles("views/home/register.html")
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	temp.Execute(w, data)
}