package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"news_web/global"
	"time"
)

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", AppConfig.Database.User, AppConfig.Database.Password, AppConfig.Database.Host, AppConfig.Database.Port, AppConfig.Database.Name)
	//fmt.Println("111111111111111111111111111:", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open database,go error :%v", err)
	}
	sqlDb, err := db.DB()

	sqlDb.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDb.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("Failed to configure database, got error: %v", err)
	}

	global.Db = db
}
