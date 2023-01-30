package userapi

import (
	"time"
	"context"
	"github.com/google/uuid"
	"github.com/kyusscaesar/go-web-app/client"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel

	Id uuid.UUID `bun:"pk,default:gen_random_uuid()"`
	Username string `bun:"notnull"`
	Email string `bun:"notnull"`
	CreatedAt time.Time `bun:"notnull,default:current_timestamp()"`
	UpdatedAt time.Time `bun:"notnull,default:current_timestamp()"`
	DeletedAt *time.Time `bun:"soft_delete"`
}

var _ bun.BeforeAppendModelHook = (*User)(nil)
func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.UpdateQuery:
		u.UpdatedAt = time.Now()
	}
	return nil
}

func (u *User) FromDTO(dto client.IUserDTO) error {
	id, err := uuid.Parse(dto.id)
	if err != nil {
		return err
	}

	u.Id = id
	u.Username = dto.Username
	u.Email = dto.Email
	u.CreatedAt = dto.CreatedAt
	u.UpdatedAt = dto.UpdatedAt
	u.DeletedAt = dto.DeletedAt
}

func (u User) IntoDTO() client.IUserDTO {
	return client.NewIUserDTO(u.Id, u.Username, u.Email, u.CreatedAt, u.UpdatedAt, u.DeletedAt)
}
