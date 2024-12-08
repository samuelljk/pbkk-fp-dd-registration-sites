package userdashboard

import (
	"log"
	"net/http"
	"pbkk-fp-dd-registration-sites/models/applicationmodel/admindashboardmodel"
	"text/template"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/user_dash.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

func GetData(w http.ResponseWriter, r *http.Request) {
	details, err := admindashboardmodel.GetAll()
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	
	data := map[string]any{
		"details": details,
	}

	// log.Printf("Template data: %+v", data)

	temp, err := template.ParseFiles("views/home/user_dash.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	temp.Execute(w, data)
}