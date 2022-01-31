package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rahul10-pu/CIGO0122/database"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetBooks(in *gin.Context) {
	books, err := database.GetBooks(h.DB) //h.DB is fully configured and can give the access of book table
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, books)
}

/*
h:=Handler{}
h.DB=GetDB()
h.GetBooks()
*/
