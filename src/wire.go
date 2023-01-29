//go:build wireinject
//+build wireinject
package main

import (
	"github.com/google/wire"
	"github.com/kyusscaesar/go-web-app/userapi"
	"github.com/kyusscaesar/go-web-app/app"
)

func GetApplication() app.Application {
	wire.Build(app.ProvideEngine, userapi.UserApiSet, app.ProvideApplication)
	return app.Application{}
}
