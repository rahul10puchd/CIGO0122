package models

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	EmailId  string
	Name     string
	Password string
	Age      int
}
