package database

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/rahul10-pu/CIGO0122/models"
)

func SignUp(db *gorm.DB, user *models.Users) error {
	//Save update value in database, if the value doesn't have primary key, will insert it
	//1. check whether email id exis or not

	//2.if email is unique then save user
	err := db.Save(user).Error //book here is without id
	if err != nil {
		return err
	}
	return nil
}
func SignIn(db *gorm.DB, user *models.Users) error {
	getUser := models.Users{}
	err := db.Select("users.*").Group("users.id").Where("users.email_id= ?", user.EmailId).First(&getUser).Error
	if err != nil {
		return err
	}
	log.Println(getUser)
	if user.Password != getUser.Password {
		return errors.New("Passowrd Incorrect")
	}
	log.Println("exiting signin")
	return nil
}
