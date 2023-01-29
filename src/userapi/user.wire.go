package userapi

import "github.com/google/wire"

func ProvideUserController() UserController {
	return NewUserController()
}

var UserApiSet = wire.NewSet(ProvideUserController)