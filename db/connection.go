package db
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)
  

func Connect(host, user, password, dbName, port string) (*gorm.DB, error) {
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:connString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	  }), &gorm.Config{})
	return db, err
}
  