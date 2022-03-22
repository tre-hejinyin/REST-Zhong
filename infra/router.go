package infra

import (
	"gorm.io/gorm"

	"server/controller"
	"server/middleware/auth"
	"server/middleware/errorhandler"
	"server/middleware/logger"
	"server/middleware/validator"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	repo "server/repository/impl"
	usecase "server/usecase/impl"
)

func SetupServer(store *gorm.DB) (*gin.Engine, error) {
	validator.Register()
	r := gin.New()
	r.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))
	r.Use(errorhandler.HandleErrors())
	r.Use(cors.Default())
	v1 := r.Group("api/v1")
	{
		userRepo := repo.NewUserRepository(store)
		userUseCase := usecase.NewUserUseCase(userRepo)
		ctrl := controller.NewUserController(userUseCase)
		v1.POST("/signup", ctrl.Signup)
		v1.POST("/signin", ctrl.Signin)
		v1.GET("/profile", auth.JWTAuthMiddleware(), ctrl.GetProfile)
		v1.POST("/profile/update", auth.JWTAuthMiddleware(), ctrl.UpdateProfile)
	}
	return r, nil
}
