package applicationcontroller

import (
	"html/template"
	"net/http"
	"pbkk-fp-dd-registration-sites/entities"
	"pbkk-fp-dd-registration-sites/models/admindashboardmodel"
	"pbkk-fp-dd-registration-sites/models/applicationmodel"
	"strconv"
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
	if r.Method == "GET" {
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

	if r.Method == "POST" {
		var user entities.User
		var university entities.University
		var degree entities.Degree

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		universityId, err := strconv.Atoi(r.FormValue("university"))
		if err != nil {
			panic(err)
		}

		degreeId, err := strconv.Atoi(r.FormValue("degree"))
		if err != nil {
			panic(err)
		}

		user.Username = r.FormValue("username")
		user.Password = r.FormValue("password")
		user.First_name = r.FormValue("first-name")
		user.Last_name = r.FormValue("last-name")
		user.Email = r.FormValue("email")
		user.Batch = r.FormValue("batch")
		university.Id = uint(universityId)
		degree.Id = uint(degreeId)


		if ok := applicationmodel.Create(user, university, degree); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/admin-dash", http.StatusSeeOther)
	}
}