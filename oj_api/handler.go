package main

import (
	"context"
	"oj_api/kitex_gen/api"
)

// ApiImpl implements the last service interface defined in the IDL.
type ApiImpl struct{}

// GetAnnouncementById implements the ApiImpl interface.
func (s *ApiImpl) GetAnnouncementById(ctx context.Context, id int) (resp *api.AnnounceResponse, err error) {
	resp = &api.AnnounceResponse{
		Title: "abc",
		UserName: "wk",
		Content: "abcdefu",
		CreateTime: "2022-04-30 12:00",
		LastUpdateTime: "2022-04-30 12:00",
	}
	return resp,nil
}
