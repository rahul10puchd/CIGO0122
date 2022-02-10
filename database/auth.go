package database

import (
	"github.com/jinzhu/gorm"
	"github.com/rahul10-pu/CIGO0122/models"
)

func SignUp(db *gorm.DB, user *models.Users) error {
	//Save update value in database, if the value doesn't have primary key, will insert it
	//1. check whether email id exis or not
	err := db.Save(user).Error //book here is without id
	if err != nil {
		return err
	}
	return nil
}
