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

// func GetDegree() []entities.Degree {
// 	rows, err := config.DB.Query("SELECT * FROM degrees")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer rows.Close()

// 	var degrees []entities.Degree

// 	for rows.Next() {
// 		var degree entities.Degree
// 		err := rows.Scan(&degree.Id, &degree.University_id, &degree.Name, &degree.Department, &degree.Duration)
// 		if err != nil {
// 			panic(err)
// 		}

// 		degrees = append(degrees, degree)
// 	}

// 	return degrees
// }

type UniversityDegree struct {
	University 	entities.University
	Degrees    	[]entities.Degree
}

func GetUniDegree() ([]UniversityDegree, error){
	rows, err := config.DB.Query(`SELECT * FROM universities JOIN degrees ON universities.university_id = degrees.university_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	universityMap := make(map[uint]UniversityDegree)

	for rows.Next() {
		var (
			uID         uint
			uName       string
			uCountry    string
			dID         uint
			dUniversityID uint
			dName       string
			dDepartment string
			dDuration   uint
		)

		err := rows.Scan(&uID, &uName, &uCountry, &dID, &dUniversityID, &dName, &dDepartment, &dDuration)
		if err != nil {
			return nil, err
		}

		degree := entities.Degree{
			Id: dID,
			University_id: dUniversityID,
			Name: dName,
			Department: dDepartment,
			Duration: dDuration,
		}

		if uniDegree, exists := universityMap[uID]; exists {
			// Append the degree to the existing slice
			uniDegree.Degrees = append(uniDegree.Degrees, degree)
			universityMap[uID] = uniDegree
		} else {
			// Create a new UniversityDegree entry
			universityMap[uID] = UniversityDegree{
				University: entities.University{
					Id:      uID,
					Name:    uName,
					Country: uCountry,
				},
				Degrees: []entities.Degree{degree},
			}
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var universityDegrees []UniversityDegree
	for _, uniDegree := range universityMap {
		universityDegrees = append(universityDegrees, uniDegree)
	}

	return universityDegrees, nil
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
			&user.Batch,
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