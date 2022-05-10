package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wannabea/OnlineJudge/oj_api/client"
	"github.com/wannabea/OnlineJudge/oj_api/kitex_gen/user"
	md "github.com/wannabea/OnlineJudge/oj_api/middleware"
)

func Update(ctx *gin.Context) {
	var info user.UserInfo
	err := ctx.BindJSON(&info)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "-1",
			"msg":    "can't find update info, err:" + err.Error(),
			"data":   nil,
		})
		return
	}

	token := ctx.Request.Header.Get("token")
	j := md.NewJWT()
	claims, err := j.ParserToken(token)
	if claims.IsAdmin == 1 {

	} else if claims.Id == info.UserId {

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "-1",
			"msg":    "have no access",
			"data":   nil,
		})
		return
	}

	log.Printf("%v", info)
	resp, err := client.UserClient.UpdateUserInfo(ctx, info.UserId,&info)

	if resp == int32(1) {
		log.Println("[INFO] update success")
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "1",
				"msg":    "update success",
				"data":   nil,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "1",
				"msg":    "update success but err:" + err.Error(),
				"data":   nil,
			})
		}
	} else {
		log.Println("[INFO] update success")
		ctx.JSON(http.StatusOK, gin.H{
			"status": "-1",
			"msg":    "update fail, err:" + err.Error(),
			"data":   nil,
		})
	}
	return
}
