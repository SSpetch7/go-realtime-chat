package db

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() *gorm.DB {
	initTimeZome()
	initConfig()

	db := initDatabase()

	AutoMigration(db)

	return db
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "-"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
func initTimeZome() {
	ict, err := time.LoadLocation("Asia/Bangkok")

	if err != nil {
		panic(err)
	}

	time.Local = ict

}

func initDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok",
		viper.GetString("db.host"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
		viper.GetString("db.port_host"),
	)

	dial := postgres.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db

}
