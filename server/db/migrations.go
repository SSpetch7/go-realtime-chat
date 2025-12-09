package db

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;type:bigint"`
	Username string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Email    string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Password string `gorm:"type:varchar(255);not null"`
}

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func CloseTable(db *gorm.DB) {
	db.Migrator().DropTable(&User{})
}
