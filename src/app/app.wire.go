package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kyusscaesar/go-web-app/userapi"
)

func ProvideEngine() *gin.Engine {
	return gin.Default()
}

func ProvideApplication(e *gin.Engine, userController userapi.UserController) Application {
	userController.ConfigureRoutes(e)

	a := Application{
		engine: e,
	}

	return a
}
