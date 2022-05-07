package main

import (
	announcement "github.com/OnlineJudge/announcement/kitex_gen/announcement/api"
	"log"
	_ "github.com/OnlineJudge/announcement/db"
)

func main() {
	svr := announcement.NewServer(new(ApiImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
