package router

import (
	_userData "alta-dashboard-be/features/users/data"
	_userHandler "alta-dashboard-be/features/users/delivery"
	_userService "alta-dashboard-be/features/users/service"
	"alta-dashboard-be/middlewares"

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
