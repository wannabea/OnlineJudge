package announcements

import (
	"context"
	"log"

	// "github.com/wannabea/OnlineJudge/oj_api/kitex_gen/api"
	"github.com/cloudwego/kitex/client"
	. "github.com/wannabea/OnlineJudge/oj_api/kitex_gen/announce"
	announce "github.com/wannabea/OnlineJudge/oj_api/kitex_gen/announce/api"
)

var c announce.Client

func init() {
	var err error
	c, err = announce.NewClient("announce", client.WithHostPorts("0.0.0.0:9092"))
	if err != nil {
		log.Fatal(err)
	}
}

func GetAnnouncementById(ctx context.Context, id int32) (resp *AnnounceResponse, err error) {
	
	// resp = &api.AnnounceResponse{
	// 	Title: "abc",
	// 	UserName: "wk",
	// 	Content: "abcdefu",
	// 	CreateTime: "2022-04-30 12:00",
	// 	LastUpdateTime: "2022-04-30 12:00",
	// }
	req := AnnounceRequest{
		AnnounceId: id,
	}
	log.Println("client to announce")
	resp, err = c.GetAnnouncementById(ctx, &req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return resp, err
}

func GetAnnouncementList(ctx context.Context) (resp []AnnouncementSimple, err error) {
	// resp = append(resp, AnnouncementSimple{
	// 	Title:          "abc",
	// 	UserName:       "wk",
	// 	CreateTime:     "2022-04-30 12:00",
	// 	LastUpdateTime: "2022-04-30 12:00",
	// })
	// resp = append(resp, AnnouncementSimple{
	// 	Title:          "abc",
	// 	UserName:       "wk",
	// 	CreateTime:     "2022-04-30 12:00",
	// 	LastUpdateTime: "2022-04-30 12:00",
	// })
	respItems, err := c.GetAllAnnouncements(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	for _, item := range respItems {
		resp = append(resp, AnnouncementSimple{
			Title:          item.Title,
			UserName:       item.UserName,
			CreateTime:     item.CreateTime,
			LastUpdateTime: item.LastUpdateTime,
		})
	}
	return resp, nil
}
