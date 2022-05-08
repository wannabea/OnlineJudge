package main

import (
	// "fmt"
	"log"
	"strconv"

	// "net"

	// "oj_api/kitex_gen/api/api"
	// "oj_api/kitex_gen/api"
	"github.com/wannabea/OnlineJudge/oj_api/handler/announcements"
	"github.com/wannabea/OnlineJudge/oj_api/handler/user"

	// "github.com/cloudwego/kitex/server"
	"github.com/gin-gonic/gin"
)

func init_router() {
	r := gin.Default()
	r.GET("/announcements", handleShowAnnouncement)
	r.GET("/username", handleGetUsername)
	r.GET("/userInfo", handleGetUserinfo)

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
