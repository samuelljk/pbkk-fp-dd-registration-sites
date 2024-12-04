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