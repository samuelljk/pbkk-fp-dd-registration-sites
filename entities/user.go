package entities

import "time"

type User struct {
	Id         uint
	Username   string
	Password   string
	Email      string
	Role       string
	First_name string
	Last_name  string
	Created_at time.Time
	Updated_at time.Time
}