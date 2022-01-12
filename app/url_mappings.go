package app

import "myapp/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.GET("/users/:user_id", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)
}
