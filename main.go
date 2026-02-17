package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"handlers/SignUp"
	"handlers/Login"
)
func main() {
	router := gin.Default()

	router.GET("/", eventhandler)

	router.POST("/signup", SignUp)
	router.POST("/login", Login)

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(auth())
	protectedRoutes.GET("/dashboard", protected)

	router.Run(":8080")
}

func eventhandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
func protected(c *gin.Context) {
	c.JSON(200, gin.H{"message": "You are authorized"})
}
