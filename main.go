package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users []User

func main() {
    router := gin.Default()

    // Routes
    router.GET("/", homeHandler)
    router.GET("/api/users", getUsersHandler)
    router.GET("/api/users/:id", getUserByIDHandler)
    router.POST("/api/users", createUserHandler)
    router.PUT("/api/users/:id", updateUserHandler)
    router.DELETE("/api/users/:id", deleteUserHandler)

    // Start server
    router.Run(":8080")
}

func homeHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Welcome to Go Web API",
        "version": "1.0.0",
    })
}

func getUsersHandler(c *gin.Context) {
    c.JSON(http.StatusOK, users)
}

func getUserByIDHandler(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
        "id": id,
        "name": "Sample User",
    })
}

func createUserHandler(c *gin.Context) {
    var user User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    users = append(users, user)
    c.JSON(http.StatusCreated, user)
}

func updateUserHandler(c *gin.Context) {
    id := c.Param("id")
    var user User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"id": id, "message": "User updated"})
}

func deleteUserHandler(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"id": id, "message": "User deleted"})
}