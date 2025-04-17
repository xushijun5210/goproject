package global

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {
	dsn := "root:xushijun5210@tcp(127.0.0.1:3306)/web_wallet_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("connect database failed, err:", err)
		return
	}
	//fmt.Println("connect database success")
	// _ = db.AutoMigrate(&model.User{})
	// fmt.Println("migrate success")
	// if err := db.AutoMigrate(&model.User{}); err != nil {
	// 	fmt.Println("create table failed, err:", err)
	// 	return
	// }
	//fmt.Println("create table success")
}
