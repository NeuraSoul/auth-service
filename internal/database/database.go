package database


import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func InitDB() *gorm.DB {
	    dsn := "host=localhost user=postgres password=yourpassword dbname=authservice port=5432 sslmode=disable"

		DB,err:= gorm.Open(postgres.Open(dsn),&gorm.Config{})
		if err != nil {
			log.Println(err)
			return nil
		}

		return DB
}
