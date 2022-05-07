package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/wannabea/OnlineJudge/announcement/kitex_gen/announce"

	. "github.com/wannabea/OnlineJudge/announcement/db"

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

func (s *ApiImpl) GetAnnouncementById(ctx context.Context, req *announce.AnnounceRequest) (resp *announce.AnnounceResponse, err error) {
	var announceItem announcementFromDb
	sqlStr := "select * from announcement where id=?"
	err = Db.QueryRow(sqlStr, 1).Scan(&announceItem.AnnounceId, &announceItem.Title, &announceItem.UserId, &announceItem.Content, &announceItem.CreateTime, &announceItem.LastUpdateTime)

	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	resp = &announce.AnnounceResponse{
		AnnounceId:     req.AnnounceId,
		Title:          announceItem.Title,
		UserName:       strconv.Itoa(announceItem.UserId),
		Content:        announceItem.Content,
		CreateTime:     strconv.Itoa(announceItem.CreateTime),
		LastUpdateTime: strconv.Itoa(announceItem.LastUpdateTime),
	}
	return
}

// GetAllAnnouncements implements the ApiImpl interface.
func (s *ApiImpl) GetAllAnnouncements(ctx context.Context) (resp []*announce.AnnounceRequest, err error) {
	// TODO: Your code here...
	return
}
