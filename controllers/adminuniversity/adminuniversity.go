package adminuniversity

import (
	"log"
	"net/http"
	"pbkk-fp-dd-registration-sites/entities"
	"pbkk-fp-dd-registration-sites/models/admindashboardmodel"
	"strconv"
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
		unidegrees, err := admindashboardmodel.GetUniDegree()
		log.Println(unidegrees)
		log.Println("X_X")
		if err != nil {
			panic(err)
		}

		data := map[string]any{
			"universities": universities,
			"unidegrees": unidegrees,
		}
	
		temp, err := template.ParseFiles("views/home/admin_uni.html")
		if err != nil {
			panic(err)
		}
	
		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		action := r.FormValue("action")

		switch action {
		case "addUniversity":
			// Extract fields for adding a new university
			var university entities.University

			university.Name = r.FormValue("name")
			university.Country = r.FormValue("country")

			// Validate required fields
			if university.Name == "" || university.Country == "" {
				http.Error(w, "Both name and country are required", http.StatusBadRequest)
				return
			}

			// Add the university
			successUni := admindashboardmodel.AddUni(university)
			if !successUni {
				http.Error(w, "Error adding university", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/admin-uni", http.StatusSeeOther)

		case "addDegree":
			// Extract fields for adding a new degree
			var degree entities.Degree

			// Expecting a university ID for associating the degree
			universityId, err := strconv.Atoi(r.FormValue("university"))
			if err != nil {
				http.Error(w, "Invalid university ID", http.StatusBadRequest)
				return
			}

			degreeDuration, err := strconv.Atoi(r.FormValue("degreeDuration"))
			if err != nil {
				http.Error(w, "Invalid degree duration", http.StatusBadRequest)
				return
			}

			degree.Name = r.FormValue("degreeName")
			degree.Department = r.FormValue("degreeDepartment")
			degree.Duration = uint(degreeDuration)
			degree.Description = r.FormValue("degreeDescription")

			// You will need to fetch the selected university 
			// to associate the new degree.
			university, err := admindashboardmodel.GetUniversityByID(universityId)
			if err != nil {
				panic(err)
			}
			if university.Id == 0 {
				http.Error(w, "Selected university not found", http.StatusBadRequest)
				return
			}

			successDegree := admindashboardmodel.AddDegree(university, degree)
			if !successDegree {
				http.Error(w, "Error adding degree", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/admin-uni", http.StatusSeeOther)

		default:
			http.Error(w, "Invalid action", http.StatusBadRequest)
		}
	}
}

// controller.go

func DeleteUni(w http.ResponseWriter, r *http.Request) {
    log.Println("DeleteUni handler triggered")
    idString := r.URL.Query().Get("id")

    id, err := strconv.Atoi(idString)
    if err != nil {
        log.Printf("Invalid ID for DeleteUni: %v", err)
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    if err := admindashboardmodel.DeleteUni(uint(id)); err != nil {
        log.Printf("Error deleting university: %v", err)
        http.Error(w, "Failed to delete university", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/admin-uni", http.StatusSeeOther)
}

func DeleteDegree(w http.ResponseWriter, r *http.Request) {
    log.Println("DeleteDegree handler triggered")
    idString := r.URL.Query().Get("id")

    id, err := strconv.Atoi(idString)
    if err != nil {
        log.Printf("Invalid ID for DeleteDegree: %v", err)
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    if err := admindashboardmodel.DeleteDeg(uint(id)); err != nil {
        log.Printf("Error deleting degree: %v", err)
        http.Error(w, "Failed to delete degree", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/admin-uni", http.StatusSeeOther)
}
