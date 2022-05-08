package main

import (
	"io"
	"log"
	"net"
	"os"

	"github.com/cloudwego/kitex/server"
	_ "github.com/wannabea/OnlineJudge/announcement/db"
	announce "github.com/wannabea/OnlineJudge/announcement/kitex_gen/announce/api"
)

func init() {
	f, err := os.OpenFile("../log/announce.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
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
	addr, _ := net.ResolveTCPAddr("tcp", ":9092")

	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := announce.NewServer(new(ApiImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
