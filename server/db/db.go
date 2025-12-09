package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	initTimeZome()
	initConfig()

	db := initDatabase()

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}
func (d *Database) GetDB() *sql.DB {
	return d.db
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

func initDatabase() *sql.DB {
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port_host"),
		viper.GetString("db.database"),
	)

	db, err := sql.Open(viper.GetString("db.driver"), dsn)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	return db

}
