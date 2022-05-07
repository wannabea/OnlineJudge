package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/wannabea/OnlineJudge/announcement/kitex_gen/announcement"

	// . "github.com/wannabea/OnlineJudge/announcement/db"

	_ "github.com/go-sql-driver/mysql"
)

// ApiImpl implements the last service interface defined in the IDL.
type ApiImpl struct{}

// GetAnnouncementById implements the ApiImpl interface.

type announcementFromDb struct {
	AnnounceId     int
	Title          string
	UserId         int
	Content        string
	CreateTime     int
	LastUpdateTime int
}

func (s *ApiImpl) GetAnnouncementById(ctx context.Context, req *announcement.AnnounceRequest) (resp *announcement.AnnounceResponse, err error) {
	var announce announcementFromDb
	sqlStr := "select * from announcement where id=?"
	err := db.QueryRow(sqlStr, 1).Scan(&announce.AnnounceId, &announce.Title, &announce.UserId, &announce.Content, &announce.CreateTime, &announce.LastUpdateTime)

	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	resp = &announcement.AnnounceResponse{
		AnnounceId:     req.AnnounceId,
		Title:          announce.Title,
		UserName:       strconv.Itoa(announce.UserId),
		Content:        announce.Content,
		CreateTime:     strconv.Itoa(announce.CreateTime),
		LastUpdateTime: strconv.Itoa(announce.LastUpdateTime),
	}
	return
}

// GetAllAnnouncements implements the ApiImpl interface.
func (s *ApiImpl) GetAllAnnouncements(ctx context.Context) (resp []*announcement.AnnounceRequest, err error) {
	// TODO: Your code here...
	return
}
