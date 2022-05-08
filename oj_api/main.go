package main

import (
	"io"
	"log"
	"os"

	// "log"
	// "net"
	// "oj_api/kitex_gen/api/api"
	// "github.com/cloudwego/kitex/server"
	// "github.com/gin-gonic/gin"
	_ "github.com/wannabea/OnlineJudge/oj_api/client"
)

func init() {
	f, err := os.OpenFile("../log/oj_api.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		f.Close()
	}()

	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	init_router()
	// addr, _ := net.ResolveTCPAddr("tcp", ":9090")

	// var opts []server.Option
	// opts = append(opts, server.WithServiceAddr(addr))

	// svr := api.NewServer(new(ApiImpl), opts...)

	// err := svr.Run()

	// if err != nil {
	// 	log.Println(err.Error())
	// }
}
