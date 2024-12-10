package applicationmodel

import (
	"pbkk-fp-dd-registration-sites/config"
	"pbkk-fp-dd-registration-sites/entities"
)

func GetUser() []entities.User{
	rows, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Role, &user.First_name, &user.Last_name, &user.Created_at, &user.Updated_at)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	return users
}

type UniversityDegree struct {
	University 	entities.University
	Degree    	entities.Degree
}


func GetUniDegreeData() ([]UniversityDegree, error){
	rows, err := config.DB.Query(`SELECT * FROM universities JOIN degrees ON universities.university_id = degrees.university_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var UniversityDegrees []UniversityDegree

	for rows.Next() {
		var university entities.University
		var degree entities.Degree
		err := rows.Scan(&university.Id, &university.Name, &university.Country, &degree.Id, &degree.University_id, &degree.Name, &degree.Department, &degree.Duration)
		if err != nil {
			return nil, err
		}
		UniversityDegrees = append(UniversityDegrees, UniversityDegree{
			University: university,
			Degree: degree,
		})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return UniversityDegrees, nil
}