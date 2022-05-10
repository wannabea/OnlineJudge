package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wannabea/OnlineJudge/oj_api/client"
	"github.com/wannabea/OnlineJudge/oj_api/kitex_gen/user"
)

func Register(ctx *gin.Context) {
	var info registerReq
	err := ctx.BindJSON(&info)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "-1",
			"msg":    "can't find register info, err:" + err.Error(),
			"data":   nil,
		})
		return
	}
	log.Printf("%v", info)
	resp, err := client.UserClient.InsertUser(ctx, &user.InsertUserInfo{
		UserName: info.UserName,
		Passwd:   info.Passwd,
		Email:    info.Email,
	})

	if resp == int32(1) {
		log.Println("[INFO] register success")
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "1",
				"msg":    "register success",
				"data":   nil,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "1",
				"msg":    "register success but err:" + err.Error(),
				"data":   nil,
			})
		}
	} else {
		log.Println("[INFO] register success")
		ctx.JSON(http.StatusOK, gin.H{
			"status": "-1",
			"msg":    "register fail, err:" + err.Error(),
			"data":   nil,
		})
	}
	return
}
