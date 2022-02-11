package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rahul10-pu/CIGO0122/database"
	"github.com/rahul10-pu/CIGO0122/models"
)

func (h *Handler) SignUp(in *gin.Context) {
	user := models.Users{}
	err := in.BindJSON(&user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.SignUp(h.DB, &user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, user)
}
func (h *Handler) SignIn(in *gin.Context) {
	user := models.Users{}
	err := in.BindJSON(&user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.SignIn(h.DB, &user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	log.Println("generating token")
	token, err := generateToken(user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	log.Println("generating token done")
	in.JSON(http.StatusOK, token)
}
func generateToken(user models.Users) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["emailid"] = user.EmailId
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString([]byte("rahul"))

	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}
func ValidateToken(Intoken string) (bool, error) {
	token, err := jwt.Parse(Intoken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("some unexpected error")
		}
		return []byte("rahul"), nil
	})
	if err != nil {
		log.Println(err)
		return false, err
	}
	if !token.Valid {
		return false, errors.New("token is not validdd")
	}
	return true, nil
}
