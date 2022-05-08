package user

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	. "github.com/wannabea/OnlineJudge/oj_api/client"
	"github.com/wannabea/OnlineJudge/oj_api/kitex_gen/user"
)

func GetUsernameById(c *gin.Context, id int32) (resp string, err error) {
	userIdStr, _ := c.Cookie("user_id")
	Address, _ := c.Cookie("address")
	IsAdminStr, _ := c.Cookie("address")
	var userId, IsAdmin int32
	fmt.Sscan(userIdStr, &userId)
	fmt.Sscan(IsAdminStr, &userId)

	idt := user.Identity{
		UserId:  userId,
		Address: Address,
		IsAdmin: IsAdmin,
	}
	resp, err = UserClient.GetNameById(c, id, &idt)
	return
}

func GetUserinfoById(c *gin.Context, id int32) (resp *user.UserInfo, err error) {
	userIdStr, _ := c.Cookie("user_id")
	log.Println("[INFO] get userIdStr from cookie :", userIdStr)
	Address, _ := c.Cookie("address")
	log.Println("[INFO] get address from cookie :", Address)
	IsAdminStr, _ := c.Cookie("IsAdminStr")
	log.Println("[INFO] get IsAdminStr from cookie :", IsAdminStr)
	var userId, IsAdmin int32
	fmt.Sscan(userIdStr, &userId)
	fmt.Sscan(IsAdminStr, &userId)
	idt := user.Identity{
		UserId:  userId,
		Address: Address,
		IsAdmin: IsAdmin,
	}
	resp, err = UserClient.GetInfoById(c, id, &idt)
	return
}
