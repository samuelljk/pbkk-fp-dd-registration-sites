package admindashboard

import (
	"net/http"
	"pbkk-fp-dd-registration-sites/models/admindashboardmodel"
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