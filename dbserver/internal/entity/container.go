package entity

type Container struct {
	Id         int    `json:"id" db:"id"`
	Port       string `json:"port" db:"port"`
	TimeOfPing string `json:"timeOfPing" db:"timeofping"`
	DateOfPing string `json:"dateOfPing" db:"dateofping"`
}

type UpdateContainers struct {
	Port       string `json:"port" db:"port"`
	TimeOfPing string `json:"timeOfPing" db:"timeofping"`
	DateOfPing string `json:"dateOfPing" db:"dateofping"`
}
