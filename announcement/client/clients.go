package client

import (
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/wannabea/OnlineJudge/announcement/kitex_gen/user/user"
)

var UserClient user.Client
var err error

func init() {
	UserClient, err = user.NewClient("user", client.WithHostPorts("0.0.0.0:9091"))
	if err != nil {
		log.Printf("[ERROR] user.newClient fail err:", err.Error())
	}
}
