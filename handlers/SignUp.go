package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthaggarwal-0001/Go.git/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context){
	var user models.User
	if err:= c.BindJSON(&user) ;err != nil{
		c.JSON(400,gin.H{"error": "Invalid user"})
		return
	}
	hashedPassword , err:=bcrypt.GenerateFromPassword([] byte(user.Password),12)
	if err!= nil{
		c.JSON(500, gin.H{"error": "Password not matched"})
		return
	}
	user.Password = string(hashedPassword)

	c.JSON(200, gin.H{"message":"Login Succesful"})

}