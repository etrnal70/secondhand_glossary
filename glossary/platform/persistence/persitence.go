package persistence

import (
	"fmt"
	"os"
	"secondhand_glossary/internal/config"
	"secondhand_glossary/internal/model"

   log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitPersistenceDB(c config.Config) *gorm.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		c.DB_USERNAME,
		c.DB_PASSWORD,
		c.DB_HOST,
		c.DB_PORT,
		c.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Println("Error opening persistence db connection: ", err)
		os.Exit(1) // TODO Handle this better, eg. using loop to wait
	}
  log.Info("Connected to mysql instance")

	// User initialization
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserRegister{})
	db.AutoMigrate(&model.UserLogin{})

	// Device initialization
	db.AutoMigrate(&model.Device{})
	db.AutoMigrate(&model.Link{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Trait{})
	db.AutoMigrate(&model.Review{})
	db.AutoMigrate(&model.Scores{})

	return db
}
