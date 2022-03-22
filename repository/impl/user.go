package impl

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"server/domain"
	"server/middleware/constants"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Store(user *domain.User) (result domain.User, err error) {
	err = u.db.Create(user).Select("first_name,last_name,email").First(&result, user.ID).Error
	if err != nil {
		err = fmt.Errorf("[userRepository.Store] failed:data=%+v   error=%w", user, err)
		return
	}
	return
}

func (u *userRepository) FindByEmail(email string) (result domain.User, err error) {
	err = u.db.First(&result, "email=?", email).Error
	if err != nil {
		err = fmt.Errorf("[userRepository.FindByEmail] failed:email=%+v error=%w", email, err)
		return
	}
	return
}

func (u *userRepository) GetProfile(ctx context.Context) (result domain.User, err error) {
	id := ctx.Value(constants.ID).(int)
	err = u.db.Select("first_name,last_name,email").First(&result, id).Error
	if err != nil {
		return domain.User{}, fmt.Errorf("[userRepository.GetProfile]failed:id=%d error=%w", id, err)
	}
	return
}

func (u *userRepository) UpdateProfile(ctx context.Context, user *domain.User) (result domain.User, err error) {
	id := ctx.Value(constants.ID).(int)
	err = u.db.Where("id=?", id).Updates(user).Select("first_name,last_name,email").First(&result, id).Error
	if err != nil {
		return domain.User{}, fmt.Errorf("[userRepository.UpdateProfile]failed:id=%d user=%+v error=%w", id, user, err)
	}
	return
}
