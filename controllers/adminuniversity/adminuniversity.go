package adminuniversity

import (
	"log"
	"net/http"
	"pbkk-fp-dd-registration-sites/entities"
	"pbkk-fp-dd-registration-sites/models/applicationmodel/admindashboardmodel"
	"text/template"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/admin_uni.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		universities := admindashboardmodel.GetUni()
		data := map[string]any{
			"universities": universities,
		}
	
		temp, err := template.ParseFiles("views/home/admin_uni.html")
		if err != nil {
			panic(err)
		}
	
		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var university entities.University

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		university.Name = r.FormValue("name")
		university.Country = r.FormValue("country")

		log.Println(university)

		if university.Name == "" || university.Country == "" {
			http.Error(w, "Both name and country are required", http.StatusBadRequest)
			return
		}
		
		success := admindashboardmodel.AddUni(university)
		if !success {
			http.Error(w, "Error adding university", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin-uni", http.StatusSeeOther)
	}
}