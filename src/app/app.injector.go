//go:build wireinject
//+build wireinject
package app

import (
	"github.com/google/wire"
	"github.com/kyusscaesar/go-web-app/userapi"
)

func GetApplication() Application {
	wire.Build(ProvideEngine, userapi.UserApiSet, ProvideApplication)
	return Application{}
}
