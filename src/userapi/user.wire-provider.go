package userapi

import "github.com/google/wire"

func ProvideUserRepository() UserRepository {
	return NewUserRepository()
}

func ProvideUserService(u *UserRepository) UserService {
	return NewUserService(u)
}

func ProvideUserController(u *UserService) UserController {
	return NewUserController(u)
}

var UserApiSet = wire.NewSet(ProvideUserService, ProvideUserController)