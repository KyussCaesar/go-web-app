package userapi

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db bun.DB
}

func NewUserRepository(db *bun.DB) UserRepository {
	return UserRepository{
		db: db
	}
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) ([]User, error) {
	users := make([]User, 0)
	err := u.db.NewSelect().Model(&users).Where("email = ?", email).Scan(ctx)
	if err != nil {
		return []User{}, err
	}

	return users, nil
}

func (u *UserRepository) Create(email string, username string) (User, error) {
	return User{}, ErrNotImplemented
}

func (u *UserRepository) FindById(id uuid.UUID) ([]User, error) {
	return []User{}, ErrNotImplemented
}

func (u *UserRepository) Update(id uuid.UUID, username string) error {
	return ErrNotImplemented
}

func (u *UserRepository) Delete(id uuid.UUID) (User, error) {
	existingUser, err := u.FindById(id)
	if err != nil {
		return User{}, err
	}

	// TODO delete and return user model
	return User{}, ErrNotImplemented
}