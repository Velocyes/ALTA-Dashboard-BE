package main

import (
	"alta-dashboard-be/app/config"
	"alta-dashboard-be/app/database"
	"alta-dashboard-be/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDB(*cfg)
	database.InitialMigration(db)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	router.InitRouter(db, e)
	
	e.Logger.Fatal(e.Start(":8081"))
}
