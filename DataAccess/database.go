package DataAccess

import (
	"fmt"
	"log"

	. "github.com/msrexe/portfolio-tracker/Core/EnvVariables"
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
	dbConfig := DSNConfig{
		User:   GoDotEnvVariable("DB_USER"),
		Pass:   GoDotEnvVariable("DB_PASS"),
		Url:    GoDotEnvVariable("DB_URL"),
		DBName: GoDotEnvVariable("DB_NAME"),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Pass, dbConfig.Url, dbConfig.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
