package admindashboard

import (
	"log"
	"net/http"
	"pbkk-fp-dd-registration-sites/models/admindashboardmodel"
	"strconv"
	"text/template"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	applicationDetails, err := admindashboardmodel.GetAll()
	if err != nil {
		panic(err)
	}

	data := map[string]any {
		"applicationDetails": applicationDetails,
	}
	

	temp, err := template.ParseFiles("views/home/admin_dash.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete handler triggered")
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := admindashboardmodel.Delete(uint(id)); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/admin-dash", http.StatusSeeOther)
}