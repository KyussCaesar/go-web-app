package app

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	ConfigureRoutes(e *gin.Engine)
}

type Application struct {
	engine *gin.Engine
	controllers []Controller
}

func (a *Application) Run() {
	a.engine.Run()
}
