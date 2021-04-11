package DataAccess

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DSNConfig struct {
	User   string
	Pass   string
	Url    string
	DBName string
}

func ConnectDB() (*gorm.DB, error) {
	dbConfig := DSNConfig{}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Pass, dbConfig.Url, dbConfig.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
