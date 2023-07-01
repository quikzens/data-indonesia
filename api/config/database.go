package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGormDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	return db
}
