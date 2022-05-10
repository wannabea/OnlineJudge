package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wannabea/OnlineJudge/oj_api/client"
	"github.com/wannabea/OnlineJudge/oj_api/kitex_gen/user"
)

type LoginReq struct {
	UserName string `json:"userName"`
	Passwd   string `json:"passwd"`
}

type registerReq struct {
	UserName string `json:"userName"`
	Passwd   string `json:"passwd"`
	Email    string `json:"email"`
}

type UserIdt struct {
	UserId   int32  `json:"userId"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	IsAdmin  int32  `json:"isAdmin"`
}

type LoginResult struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

func LoginCheck(ctx *gin.Context, req LoginReq) (isPass bool, idt UserIdt, err error) {
	id, err := client.UserClient.GetUserIdByUserName(ctx, req.UserName)
	if id == -1 {
		log.Println("[ERROR] can't find the user name")
		return
	}
	var info *user.UserInfo
	info, err = client.UserClient.GetInfoById(ctx, id)
	if err != nil {
		return
	}
	if info.Passwd != req.Passwd {
		log.Println("[ERROR] password incorrect")
		err = PasswdIncorrect
		return
	}
	idt.IsAdmin = info.IsAdmin
	idt.UserId = info.UserId
	idt.UserName = info.UserName
	idt.Email = info.Email
	isPass = true
	return

}
