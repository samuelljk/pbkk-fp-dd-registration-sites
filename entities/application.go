package entities

type Application struct {
	Id            uint
	User_id       uint
	University_id uint
	Degree_id     uint
	Status        string
	Submitted_at  []uint8
	Updated_at    []uint8
}