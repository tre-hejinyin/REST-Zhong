package repository

import (
	"context"

	"server/domain"
)

type UserRepository interface {
	Store(user *domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	GetProfile(ctx context.Context) (domain.User, error)
	UpdateProfile(ctx context.Context, user *domain.User) (domain.User, error)
}
