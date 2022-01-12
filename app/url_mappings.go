package app

import "myapp/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
