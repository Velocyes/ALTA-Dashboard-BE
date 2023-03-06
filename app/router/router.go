package router

import (
	_userData "alta-dashboard-be/features/users/data"
	_userHandler "alta-dashboard-be/features/users/delivery"
	_userService "alta-dashboard-be/features/users/service"
	_logData "alta-dashboard-be/features/logs/data"
	_logHandler "alta-dashboard-be/features/logs/delivery"
	_logService "alta-dashboard-be/features/logs/service"
	"alta-dashboard-be/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initUserRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandler := _userHandler.New(userService)

	e.GET("/users", userHandler.GetAllUser)
	e.GET("/users/:user_id", userHandler.GetUserData, middlewares.JWTMiddleware())
	e.POST("/users/login", userHandler.Login)
	e.POST("/users", userHandler.Register)
	// e.POST("/users", userHandler.Register, middlewares.JWTMiddleware())
	e.PUT("/users/:user_id", userHandler.UpdateAccount, middlewares.JWTMiddleware())
	e.DELETE("/users/:user_id", userHandler.RemoveAccount, middlewares.JWTMiddleware())
}

func initLogRouter(db *gorm.DB, e *echo.Echo) {
	logData := _logData.New(db)
	logService := _logService.New(logData)
	logHandler := _logHandler.New(logService)

	e.POST("/logs", logHandler.AddLog, middlewares.JWTMiddleware())
	e.GET("/mentees/:mentee_id/logs", logHandler.GetLogDataByMenteeId)
	// e.POST("/users/login", userHandler.Login)
	// e.POST("/users", userHandler.Register, middlewares.JWTMiddleware())
	// e.PUT("/users/:user_id", userHandler.UpdateAccount, middlewares.JWTMiddleware())
	// e.DELETE("/users/:user_id", userHandler.RemoveAccount, middlewares.JWTMiddleware())
}

func InitRouter(db *gorm.DB, e *echo.Echo) {
	initUserRouter(db, e)
	initLogRouter(db, e)
}
