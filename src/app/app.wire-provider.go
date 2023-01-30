package app

import (
	"github.com/google/wire"
	"github.com/gin-gonic/gin"
	"github.com/kyusscaesar/go-web-app/monitoringapi"
	"github.com/kyusscaesar/go-web-app/userapi"
)

func ProvideEngine() *gin.Engine {
	e := gin.Default()

	e.SetTrustedProxies(nil)

	return e
}

func ProvideApplication(e *gin.Engine, monitoringController monitoringapi.MonitoringController, userController userapi.UserController) Application {
  monitoringController.ConfigureRoutes(e)
	userController.ConfigureRoutes(e)

	a := Application{
		engine: e,
	}

	return a
}

var AppSet = wire.NewSet(ProvideEngine, ProvideApplication)