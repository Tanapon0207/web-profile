package data


import (
	"gorm.io/gorm"
)
type User struct {
	gorm.Model
	Fullname string
	Email    string
	Username string
	Password string
	Isuse    string

	// Role     string
}
