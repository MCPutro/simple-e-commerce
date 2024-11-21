package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlConnection() (*sql.DB, error) {
	var db *sql.DB
	var err error
	maxRetry := 10

	for i := 0; i < maxRetry; i++ {
		db, err = sql.Open("mysql", "mcputro:welcome1@tcp(localhost:1123)/e_commerce?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true")
		if err != nil {
			fmt.Println("error :", err)
		}

		if i > 0 {
			fmt.Println("DB Connection : Retry Mechanism [" + strconv.Itoa(i) + "x]")
		}
		if err == nil {
			db.SetMaxOpenConns(25)
			db.SetMaxIdleConns(5)
			db.SetConnMaxLifetime(5 * time.Minute)
			break
		}
		time.Sleep(5 * time.Second)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error connecting to database: ", err)
		return nil, err
	}

	return db, nil
}
