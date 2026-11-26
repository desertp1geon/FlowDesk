package users

import (
	"gorm.io/gorm"
)

const (
	OWNER = iota
	USER
	CLIENT
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	PasswordHash string
	Role         int
}
