package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"alta-dashboard-be/app/config"
	_logModel "alta-dashboard-be/features/logs/models"
	_menteeModel "alta-dashboard-be/features/mentee/models"
	_classModel "alta-dashboard-be/features/class/models"
	_userModel "alta-dashboard-be/features/users/models"
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
	userGorm := _userModel.User{Role: consts.E_USER_Admin}
	db.Model(userGorm).First(&userGorm)
	if userGorm.ID == 0 {
		hashedPassword, _ := _userService.HashPassword("qwerty")
		db.Model(userGorm).Save(&_userModel.User{FullName: "Admin", Email: "Admin@gmail.com", Password: hashedPassword, Role: consts.E_USER_Admin})
	}
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_userModel.User{}, _classModel.Class{}, _menteeModel.Emergency{}, _menteeModel.Education{}, &_menteeModel.Mentee{}, &_logModel.Log{})
	initSuperAdmin(db)
}
