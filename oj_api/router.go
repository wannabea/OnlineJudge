package main

import (
	// "fmt"
	"log"
	"net/http"
	"strconv"

	// "net"

	"github.com/wannabea/OnlineJudge/oj_api/handler/announcements"
	"github.com/wannabea/OnlineJudge/oj_api/handler/user"
	md "github.com/wannabea/OnlineJudge/oj_api/middleware"

	// "github.com/cloudwego/kitex/server"
	"github.com/gin-gonic/gin"
)

func GetDataByTime(c *gin.Context) {
	// 上面我们在JWTAuth()中间中将'claims'写入到gin.Context的指针对象中，因此在这里可以将之解析出来
	claims := c.MustGet("claims").(*md.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}

func init_router() {
	r := gin.Default()
	r.GET("/announcements", handleShowAnnouncement)
	r.GET("/username", handleGetUsername)
	r.GET("/userInfo", handleGetUserinfo)
	r.GET("/userId", handleGetUserId)

	r.POST("/login", user.Login)
	r.POST("/register", user.Register)

	sv := r.Group("/auth")
	sv.Use(md.JWTAuth())

	{
		sv.GET("/test", GetDataByTime)
		sv.POST("/update", user.Update)
	}

	err := r.Run(":9090")

	if err != nil {
		log.Println(err.Error())
	}
}

func handleShowAnnouncement(c *gin.Context) {
	idString, ok := c.GetQuery("id")
	id, _ := strconv.Atoi(idString)
	if ok {
		announcement, err := announcements.GetAnnouncementById(c, int32(id))
		if err == nil {
			c.JSON(200, announcement)
		} else {
			c.String(404, "Not Found")
		}
	} else {
		announcementList, err := announcements.GetAnnouncementList(c)
		if err == nil {
			c.JSON(200, announcementList)
			log.Println("resp :", announcementList)
		} else {
			c.String(404, "Not Found")
		}
	}
}

func handleGetUsername(c *gin.Context) {
	idString, ok := c.GetQuery("id")
	id, _ := strconv.Atoi(idString)
	if ok {
		name, err := user.GetUsernameById(c, int32(id))
		if err == nil {
			c.String(200, name)
		} else {
			c.String(404, "Not Found")
		}
	} else {
		c.String(404, "Not Found")
	}
}

func handleGetUserinfo(c *gin.Context) {
	idString, ok := c.GetQuery("id")
	id, _ := strconv.Atoi(idString)
	if ok {
		info, err := user.GetUserinfoById(c, int32(id))
		if err == nil {
			c.JSON(200, info)
		} else {
			c.String(404, "Not Found")
		}
	} else {
		c.String(404, "Not Found")
	}
}

func handleGetUserId(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if ok {
		id, err := user.GetUserIdByName(c, name)
		if err == nil {
			c.JSON(200, id)
		} else {
			c.String(404, "Not Found")
		}
	} else {
		c.String(404, "Not Found")
	}
}
