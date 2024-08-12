package model

import "gorm.io/gorm"

type Password struct {
	gorm.Model
	Name     string
	Username string
	Password string
}
