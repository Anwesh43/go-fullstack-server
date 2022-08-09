package db 

import (
	"fullstackserver/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, tableName string) {
	tableMap := make(map[string]interface{})
	tableMap["author"] = &models.Author{}
	tableMap["book"] = &models.Book{}
	db.AutoMigrate(tableMap[tableName])
}

func Create(db *gorm.DB, obj interface{}) {
	db.Create(obj)
}

func QueryAll(db *gorm.DB, obj interface{}) interface{} {
	db.Find(obj)
	return obj 
}
