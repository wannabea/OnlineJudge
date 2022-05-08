package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type announcementFromDb struct {
	AnnounceId     int
	Title          string
	UserId         int
	Content        string
	CreateTime     int
	LastUpdateTime int
}

// var Db *sql.DB

func test() (err error) {
	dsn := "oj:onlinejudge2022@tcp(127.0.0.1:3306)/onlinejudge?charset=utf8mb4&parseTime=True"

	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("[ERROR] open onlinejudge database fail")
		log.Println(err.Error())
		return err
	}
	log.Println("[INFO] open onlinejudge database success")

	err = Db.Ping()
	if err != nil {
		return err
	}

	var announceItem announcementFromDb
	sqlStr := "select * from announcements where announce_id=?"
	err = Db.QueryRow(sqlStr, 1).Scan(&announceItem.AnnounceId, &announceItem.Title, &announceItem.UserId, &announceItem.Content, &announceItem.CreateTime, &announceItem.LastUpdateTime)

	if err != nil {
		log.Printf("scan failed, err:%v\n", err)
		return
	}
	log.Print(announceItem)
	return err
}

func main() {
	err := test()
	if err != nil {
		log.Println(err.Error())
	}
}
