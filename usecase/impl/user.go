package impl

import (
	"context"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"server/domain"
	"server/middleware/cache"
	"server/middleware/errorhandler"
	"server/middleware/jwt"
	"server/repository"
)

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *userUseCase {
	return &userUseCase{repo: repo}
}

func (u *userUseCase) Signup(request *domain.SignupRequest) (domain.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, fmt.Errorf("[userUseCase.Signup] hash password failed:password=%s,error=%w",
			request.Password, err)
	}
	result, err := u.repo.Store(&domain.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  string(password),
	})
	if err != nil {
		return domain.User{}, fmt.Errorf("[userUseCase.Signup]create failed:error=%w", err)
	}
	return result, nil
}

func (u *userUseCase) Signin(request *domain.SigninRequest) (string, error) {
	user, err := u.repo.FindByEmail(request.Email)
	if err != nil {
		return "", fmt.Errorf("[userUseCase.Signin] find by email failed:error=%w", err)
	}
	// verify password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return "", fmt.Errorf("[userUseCase.Signin] password invalid:request=%+v err=%w", request, errorhandler.ErrorEmailOrPassword)
	}
	// generate token
	token, err := jwt.GenToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("[userUseCase.Signin]gen token failed: error=%w", err)
	}
	// save to redis
	err = cache.Set(strconv.Itoa(user.ID), token)
	if err != nil {
		return "", fmt.Errorf("[userUseCase.Signin]cache hashset failed: error=%w", err)
	}
	return token, err
}

func (u *userUseCase) FetchProfile(ctx context.Context) (domain.User, error) {
	result, err := u.repo.GetProfile(ctx)
	if err != nil {
		return domain.User{}, fmt.Errorf("[userUseCase.FetchProfile]failed:error=%w", err)
	}
	return result, nil
}

func (u *userUseCase) UpdateProfile(ctx context.Context, request *domain.UpdateProfileRequest) (domain.User, error) {
	result, err := u.repo.UpdateProfile(ctx, &domain.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
	})
	if err != nil {
		return domain.User{}, fmt.Errorf("[userUseCase.UpdateProfile]update failed:error=%w", err)
	}
	return result, nil
}
