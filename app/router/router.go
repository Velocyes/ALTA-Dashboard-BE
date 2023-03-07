package router

import (
	_classData "alta-dashboard-be/features/class/data"
	_classHandler "alta-dashboard-be/features/class/delivery"
	_classService "alta-dashboard-be/features/class/service"
	_logData "alta-dashboard-be/features/logs/data"
	_logHandler "alta-dashboard-be/features/logs/delivery"
	_logService "alta-dashboard-be/features/logs/service"
	_menteeData "alta-dashboard-be/features/mentee/data"
	_menteeHandler "alta-dashboard-be/features/mentee/delivery"
	_menteeService "alta-dashboard-be/features/mentee/service"
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
	e.GET("/users/:user_id", userHandler.GetUserData, middlewares.JWTMiddleware())
	e.POST("/users/login", userHandler.Login)
	e.POST("/users", userHandler.Register, middlewares.JWTMiddleware())
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

func initClassRouter(db *gorm.DB, e *echo.Echo) {
	classData := _classData.New(db)
	classService := _classService.New(classData)
	classHandler := _classHandler.New(classService)

	e.GET("/classes", classHandler.GetAll)
	e.GET("/classes/:id", classHandler.GetOne, middlewares.JWTMiddleware())
	e.POST("/classes", classHandler.Create, middlewares.JWTMiddleware())
	e.PUT("/classes/:id", classHandler.Update, middlewares.JWTMiddleware())
	e.DELETE("/classes/:id", classHandler.Delete, middlewares.JWTMiddleware())
}

func initMenteeRouter(db *gorm.DB, e *echo.Echo) {
	menteeData := _menteeData.New(db)
	menteeService := _menteeService.New(menteeData)
	menteeHandler := _menteeHandler.New(menteeService)

	e.GET("/mentees", menteeHandler.GetAll)
	e.GET("/mentees/:id", menteeHandler.GetOne, middlewares.JWTMiddleware())
	e.POST("/mentees", menteeHandler.Create, middlewares.JWTMiddleware())
	e.PUT("/mentees/:id", menteeHandler.Update, middlewares.JWTMiddleware())
	e.DELETE("/mentees/:id", menteeHandler.Delete, middlewares.JWTMiddleware())
}

func InitRouter(db *gorm.DB, e *echo.Echo) {
	initUserRouter(db, e)
	initClassRouter(db, e)
	initMenteeRouter(db, e)
	initLogRouter(db, e)
}
