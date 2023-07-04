package user

import (
	"context"
	"time"
)

type User struct {
	ID        string     `json:"id" db:"user_id"`
	FirstName string     `json:"first_name" db:"first_name"`
	LastName  string     `json:"last_name" db:"last_name"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"-" db:"-"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type CreateUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type CreateUserRes struct {
	ID   string `json:"id" db:"user_id"`
	Role string `json:"role" db:"user_role"`
}

type LoginUserReq struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"-" db:"-"`
}

type LoginUserRes struct {
	ID           string `json:"id" db:"user_id"`
	Role         string `json:"role" db:"user_role"`
	AccessToken  string `json:"access_token" db:"-"`
	RefreshToken string `json:"refresh_token" db:"-"`
}

type Service interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error)
	Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error)
	Logout(c context.Context, userToken string) error
}
