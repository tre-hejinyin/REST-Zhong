package usecase

import (
	"context"

	"server/domain"
)

type UserUseCase interface {
	Signup(user *domain.SignupRequest) (domain.User, error)
	Signin(user *domain.SigninRequest) (string, error)
	FetchProfile(ctx context.Context) (domain.User, error)
	UpdateProfile(ctx context.Context, request *domain.UpdateProfileRequest) (domain.User, error)
}
