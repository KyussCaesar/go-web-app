package app

import (
	"github.com/gin-gonic/gin"
)

type Application struct {
	engine *gin.Engine
}

func (a *Application) Run() {
	a.engine.Run()
}
