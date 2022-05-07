package announcements

import (
	// "log"
	"context"
	// "oj_api/kitex_gen/api"
	// "github.com/gin-gonic/gin"
)

func GetAnnouncementById(ctx context.Context, id int) (resp Announcement, err error) {
	resp = Announcement{
		Title: "abc",
		UserName: "wk",
		Content: "abcdefu",
		CreateTime: "2022-04-30 12:00",
		LastUpdateTime: "2022-04-30 12:00",
	}
	return resp,  nil
}

func GetAnnouncementList(ctx context.Context) (resp []AnnouncementSimple, err error) {
	resp = append(resp, AnnouncementSimple{
		Title: "abc",
		UserName: "wk",
		CreateTime: "2022-04-30 12:00",
		LastUpdateTime: "2022-04-30 12:00",
	})
	resp = append(resp, AnnouncementSimple{
		Title: "abc",
		UserName: "wk",
		CreateTime: "2022-04-30 12:00",
		LastUpdateTime: "2022-04-30 12:00",
	})
	return resp,  nil
}

