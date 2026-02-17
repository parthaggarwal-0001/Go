func authrouter() {
	router := gin.Default()

	router.GET("/", eventhandler)

	router.POST("/signup", signup)
	router.POST("/login", login)

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(AuthMiddleware())
	protectedRoutes.GET("/dashboard", protected)

	router.Run(":8080")
}
