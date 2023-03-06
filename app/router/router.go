package router

import (
	_userData "immersive-dashboard-app/features/users/data"
	_userHandler "immersive-dashboard-app/features/users/delivery"
	_userService "immersive-dashboard-app/features/users/service"
	"immersive-dashboard-app/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initUserRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandler := _userHandler.New(userService)

	e.GET("/users", userHandler.GetAllUser)
	e.GET("/users/:id", userHandler.GetUserData, middlewares.JWTMiddleware())
	e.POST("/users/login", userHandler.Login)
	e.POST("/users", userHandler.Register, middlewares.JWTMiddleware())
	e.PUT("/users/:id", userHandler.UpdateAccount, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userHandler.RemoveAccount, middlewares.JWTMiddleware())
}

func InitRouter(db *gorm.DB, e *echo.Echo) {
	initUserRouter(db, e)
}
