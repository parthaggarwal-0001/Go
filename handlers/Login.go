package handlers

import (
	"go/token"
	"net/http"
	"github.com/gin-gonic/gin"
)
func Login(c *gin.Context){
	var user User
	var storedUser User

	if err:= c.BindJSON(&user) ; err!= nil{
		c.JSON(400,gin.H{"error":"Invalid User"})
		return
	}
	storedUser User{
		UserName: "testing"
		Password: "123456"
	}
	err.bycrpt.compareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err!=nil{
		c.JSON(400, gin.H{"error":"Invaild Credientials"})
		return 
	}
	token, _ := gernerateToken(user.Password)
	c.JSON(200, gin.H{"token",token})
}