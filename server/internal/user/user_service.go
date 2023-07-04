package user

import (
	"context"
	"time"

	"github.com/nedpals/supabase-go"
	"github.com/rs/zerolog"
)

type service struct {
	client  *supabase.Client
	timeout time.Duration
	log     zerolog.Logger
}

func NewService(client *supabase.Client, l zerolog.Logger) Service {
	return &service{
		client:  client,
		timeout: time.Duration(30) * time.Second,
		log:     l,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	result, err := s.client.Auth.SignUp(ctx, supabase.UserCredentials{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		s.log.Error().Err(err).Msg("Error creating user")
		return nil, err
	}

	res := &CreateUserRes{
		ID:   result.ID,
		Role: result.Role,
	}

	return res, nil
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	result, err := s.client.Auth.SignIn(ctx, supabase.UserCredentials{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		s.log.Error().Err(err).Msg("Error logging in")
		return &LoginUserRes{}, nil
	}

	return &LoginUserRes{
		ID:           result.User.ID,
		Role:         result.User.Role,
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	}, nil
}

func (s *service) Logout(c context.Context, userToken string) error {
	if err := s.client.Auth.SignOut(c, userToken); err != nil {
		return err
	}

	return nil
}
