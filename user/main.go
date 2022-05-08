package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	_ "github.com/wannabea/OnlineJudge/user/db"
	"github.com/wannabea/OnlineJudge/user/kitex_gen/user/user"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":9091")

	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := user.NewServer(new(UserImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
