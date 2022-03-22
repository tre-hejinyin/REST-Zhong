package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"server/domain"
	"server/middleware/errorhandler"
	"server/middleware/response"
	"server/usecase"
)

type user struct {
	usecase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) *user {
	return &user{usecase: usecase}
}

func (u *user) Signup(c *gin.Context) {
	var request domain.SignupRequest
	err := c.ShouldBind(&request)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	result, err := u.usecase.Signup(&request)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, response.Body{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: result,
	})
}

func (u *user) Signin(c *gin.Context) {
	var request domain.SigninRequest
	err := c.ShouldBind(&request)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	result, err := u.usecase.Signin(&request)
	if err != nil {
		if errors.Is(err, errorhandler.ErrorEmailOrPassword) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.Body{
				Code: errorhandler.CodeEmailOrPasswordError,
				Msg:  err.Error(),
			})
			return
		}
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, response.Body{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: result,
	})
}

func (u *user) GetProfile(c *gin.Context) {
	result, err := u.usecase.FetchProfile(c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, response.Body{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: result,
	})
}

func (u *user) UpdateProfile(c *gin.Context) {
	var request domain.UpdateProfileRequest
	err := c.ShouldBind(&request)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	result, err := u.usecase.UpdateProfile(c, &request)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, response.Body{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: result,
	})
}
