package database

import (
	"github.com/jinzhu/gorm"
	"github.com/rahul10-pu/CIGO0122/models"
)

//GetBooks is creating connection and interacting from golang app to db server via db variable
func GetBooks(db *gorm.DB) ([]models.Book, error) {

	books := []models.Book{}
	query := db.Select("books.*")
	err := query.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
func GetBookByID(db *gorm.DB, id string) (models.Book, error) {
	return models.Book{}, nil
}
func DeleteBookByID(db *gorm.DB, id string) error {
	return nil
}
func UpdateBookByID(db *gorm.DB, book *models.Book) error {
	return nil
}
