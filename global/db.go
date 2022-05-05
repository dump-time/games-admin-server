package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lixiao189/games-admin-server/log"
)

// DB is global database object
var DB *gorm.DB

// init database connection by configurations
func initDB() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.DB.Username,
		Config.DB.Password,
		Config.DB.Hostname,
		Config.DB.Port,
		Config.DB.DBName,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
