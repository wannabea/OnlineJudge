package db

import (
	"log"
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "oj:onlinejudge2022@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	
	db, err = sql.Open("onlinejudge", dsn)
	if err != nil {
		return err
	}
	log.Println("db login")

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
}