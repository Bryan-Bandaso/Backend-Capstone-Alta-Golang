package auth

import (
	auths "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/controller/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	m "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/middleware"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/repository/admin"
	auth2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/service/admin"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := admin.NewAuthRepository(db)
	svc := auth2.NewAuthService(repo, conf)

	controller := auths.AuthController{
		Svc: svc,
	}

	echo.POST("/admin/register", controller.Register)
	echo.POST("/admin/login", controller.Login)
	echo.GET("/admin/user/:name", controller.GetUser, m.JWTTokenMiddleware())
}
