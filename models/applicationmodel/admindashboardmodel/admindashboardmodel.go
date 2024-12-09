package admindashboardmodel

import (
	"pbkk-fp-dd-registration-sites/config"
	"pbkk-fp-dd-registration-sites/entities"
)

func GetUni() []entities.University {
	rows, err := config.DB.Query("SELECT * FROM universities")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var universities []entities.University

	for rows.Next() {
		var university entities.University
		err := rows.Scan(&university.Id, &university.Name, &university.Country)
		if err != nil {
			panic(err)
		}

		universities = append(universities, university)
	}

	return universities
}

func AddUni(university entities.University) bool {
	result, err := config.DB.Exec(`
		INSERT INTO universities (name, country) 
		VALUE (?, ?)`,
		university.Name,
		university.Country,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

type ApplicationDetails struct {
	Application entities.Application
	User        entities.User
	University  entities.University
	Degree      entities.Degree
}

func GetAll() ([]ApplicationDetails, error) {
	rows, err := config.DB.Query(`
		SELECT applications.*, users.*, universities.*, degrees.*
		FROM applications
		JOIN users ON applications.user_id = users.user_id
		JOIN universities ON applications.university_id = universities.university_id
		JOIN degrees ON applications.degree_id = degrees.degree_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var applicationDetails []ApplicationDetails
	
	for rows.Next() {
		var application entities.Application
		var user entities.User
		var university entities.University
		var degree entities.Degree
		err := rows.Scan(
			&application.Id,
			&application.User_id,
			&application.University_id,
			&application.Degree_id,
			&application.Status,
			&application.Submitted_at,
			&application.Updated_at,
			&user.Id,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.Role,
			&user.First_name,
			&user.Last_name,
			&user.Created_at,
			&user.Updated_at,
			&university.Id,
			&university.Name,
			&university.Country,
			&degree.Id,
			&degree.University_id,
			&degree.Name,
			&degree.Department,
			&degree.Duration,
		)
		if err != nil {
			return nil, err
		}
		applicationDetails = append(applicationDetails, ApplicationDetails{
			Application: application,
			User:        user,
			University:  university,
			Degree:      degree,
		})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return applicationDetails, nil
}