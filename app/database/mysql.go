package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"alta-dashboard-be/app/config"
	_classMdl "alta-dashboard-be/features/class/models"
	_menteeMdl "alta-dashboard-be/features/mentee/models"
	_userData "alta-dashboard-be/features/users/data"
)

func InitDB(cfg config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Println("error connect to DB", err.Error())
		return nil
	}

	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_userData.User{}, &_classMdl.Class{}, &_menteeMdl.Emergency{}, &_menteeMdl.Education{}, &_menteeMdl.Mentee{})
}
