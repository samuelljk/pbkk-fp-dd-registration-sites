package admindashboard

import (
	"net/http"
	"text/template"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/admin_dash.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}