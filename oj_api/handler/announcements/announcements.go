package announcements

import (
	"context"
	"log"

	// "github.com/wannabea/OnlineJudge/oj_api/kitex_gen/api"

	. "github.com/wannabea/OnlineJudge/oj_api/client"
	. "github.com/wannabea/OnlineJudge/oj_api/kitex_gen/announce"
)

func GetAnnouncementById(ctx context.Context, id int32) (resp *AnnounceInfo, err error) {

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

	// AnnounceClient, err = announce.NewClient("announce", client.WithHostPorts("0.0.0.0:9092"))
	log.Println("client to announce")
	resp, err = AnnounceClient.GetAnnouncementById(ctx, &req)
	log.Println("get resp to announce")
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
	respItems, err := AnnounceClient.GetAllAnnouncements(ctx)
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
