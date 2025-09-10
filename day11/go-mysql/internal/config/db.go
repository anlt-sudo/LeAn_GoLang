package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB{
	LoadEnv()

	cfg := mysql.NewConfig()
	cfg.User = GetEnv("DB_USER", "root")
	cfg.Passwd = GetEnv("DB_PASS", "")
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", GetEnv("DB_HOST", "127.0.0.1"), GetEnv("DB_PORT", "3306"))
	cfg.DBName = GetEnv("DB_NAME", "album_management")

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil{
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil{
		log.Fatal(err)
	}

	fmt.Println("Connect DB Successfully")

	return db
}