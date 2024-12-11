package entities

type User struct {
	Id         uint
	Username   string
	Password   string
	Email      string
	Batch      string
	Role       string
	First_name string
	Last_name  string
	Created_at []uint8
	Updated_at []uint8
}