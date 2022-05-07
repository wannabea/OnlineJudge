package main

import (
	announcement "github.com/wannabea/OnlineJudge/announcement/kitex_gen/announcement/api"
	"log"
	// _ "github.com/wannabea/OnlineJudge/announcement/db"
)

func main() {
	svr := announcement.NewServer(new(ApiImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
