package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	_ "github.com/wannabea/OnlineJudge/announcement/db"
	announce "github.com/wannabea/OnlineJudge/announcement/kitex_gen/announce/api"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":9092")

	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := announce.NewServer(new(ApiImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
