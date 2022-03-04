package api

import (
	"os"

	controllers "github.com/gabrielnando1/go_service_transfer/src/presenter/api/controllers"
	//configuration "graduel.dentist.api/src/app/api/configs"
)

var (
	router = controllers.Router
)

func StartApi() {

	controllers.Initialize()
	//router.GET("/user/signin", controllers.UserController_SignIn)

	port := os.Getenv("PORT") //using heroku host
	if port == "" {
		port = "8000" //localhost
	}

	router.Run(":" + port)
}
