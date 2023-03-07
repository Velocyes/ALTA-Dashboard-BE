package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"alta-dashboard-be/app/config"
	_logData "alta-dashboard-be/features/logs/data"
	_userData "alta-dashboard-be/features/users/data"
	_userService "alta-dashboard-be/features/users/service"
	"alta-dashboard-be/utils/consts"
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

func initSuperAdmin(db *gorm.DB) {
	userGorm := _userData.User{Role: consts.E_USER_Admin}
	db.Model(userGorm).First(&userGorm)
	if userGorm.ID == 0 {
		hashedPassword, _ := _userService.HashPassword("qwerty")
		db.Model(userGorm).Save(&_userData.User{FullName: "User", Email: "User@gmail.com", Password: hashedPassword, Role: consts.E_USER_Admin})
	}
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_userData.User{}, &_logData.Mentee{}, &_logData.Log{})
	initSuperAdmin(db)
}
