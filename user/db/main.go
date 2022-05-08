package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var Db *sql.DB

func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "oj:onlinejudge2022@tcp(127.0.0.1:3306)/onlinejudge?charset=utf8mb4&parseTime=True"

	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("[ERROR] open onlinejudge database fail")
		return err
	}
	log.Println("[INFO] open onlinejudge database success")

	err = Db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
}
