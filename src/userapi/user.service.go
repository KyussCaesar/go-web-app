package userapi

import "github.com/kyusscaesar/go-web-app/client"

type UserService struct {
	userRepository UserRepository
}

func NewUserService(u *UserRepository) UserService {
	return UserService{
		userRepository: u
	}
}

func (u *UserService) CreateUser(ctx context.Context, email string, username string) (User, error) {
	existingUser, err := u.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}

	if len(existingUser) != 0 {
		return User{}, ErrUserAlreadyExists
	}

	user, err := u.userRepository.Create(email, username)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
