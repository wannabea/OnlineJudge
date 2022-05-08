package client

import (
	"log"

	"github.com/cloudwego/kitex/client"
	announce "github.com/wannabea/OnlineJudge/oj_api/kitex_gen/announce/api"
	"github.com/wannabea/OnlineJudge/oj_api/kitex_gen/user/user"
)

var UserClient user.Client
var AnnounceClient announce.Client
var err error

func init() {
	UserClient, err = user.NewClient("user", client.WithHostPorts("0.0.0.0:9091"))
	if err != nil {
		log.Printf("[ERROR] user.newClient fail err:", err.Error())
	}
	AnnounceClient, err = announce.NewClient("announce", client.WithHostPorts("0.0.0.0:9092"))
	if err != nil {
		log.Printf("[ERROR] user.newClient fail err:", err.Error())
	}
	log.Println("[INFO] init clients success")
}
