package main

import (
	"context"
	"log"

	// "github.com/wannabea/OnlineJudge/oj_api/kitex_gen/api"
	"github.com/cloudwego/kitex/client"
	. "github.com/wannabea/OnlineJudge/oj_api/kitex_gen/announce"
	announce "github.com/wannabea/OnlineJudge/oj_api/kitex_gen/announce/api"
)

// ApiImpl implements the last service interface defined in the IDL.
type ApiImpl struct{}

func (s *ApiImpl) GetAnnouncementById(ctx context.Context, id int32) (resp *AnnounceResponse, err error) {
	c, err := announce.NewClient("announce", client.WithHostPorts("0.0.0.0:9091"))
	if err != nil {
		log.Fatal(err)
	}
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
	return resp, err
}
