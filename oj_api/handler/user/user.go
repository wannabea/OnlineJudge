package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/wannabea/OnlineJudge/oj_api/client"
	"github.com/wannabea/OnlineJudge/oj_api/kitex_gen/user"
)

func GetUsernameById(c *gin.Context, id int32) (resp string, err error) {
	resp, err = UserClient.GetNameById(c, id)
	return
}

func GetUserinfoById(c *gin.Context, id int32) (resp *user.UserInfo, err error) {
	resp, err = UserClient.GetInfoById(c, id)
	return
}

func GetUserIdByName(c *gin.Context, name string) (resp int32, err error) {
	resp, err = UserClient.GetUserIdByUserName(c, name)
	return
}
