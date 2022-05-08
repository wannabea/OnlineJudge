package main

import (
	"context"
	"log"
	"time"

	"github.com/wannabea/OnlineJudge/announcement/kitex_gen/announce"

	. "github.com/wannabea/OnlineJudge/announcement/client"
	. "github.com/wannabea/OnlineJudge/announcement/db"

	_ "github.com/go-sql-driver/mysql"
)

// ApiImpl implements the last service interface defined in the IDL.
type ApiImpl struct{}

// GetAnnouncementById implements the ApiImpl interface.

type announcementFromDb struct {
	AnnounceId     int32
	Title          string
	UserId         int32
	Content        string
	CreateTime     int64
	LastUpdateTime int64
	visible        int
}

func (s *ApiImpl) GetAnnouncementById(ctx context.Context, req *announce.AnnounceRequest) (resp *announce.AnnounceResponse, err error) {
	var announceItem announcementFromDb
	sqlStr := "select * from announcements where announce_id=? and visible=1"
	err = Db.QueryRow(sqlStr, req.AnnounceId).Scan(&announceItem.AnnounceId, &announceItem.Title, &announceItem.UserId, &announceItem.Content, &announceItem.CreateTime, &announceItem.LastUpdateTime, &announceItem.visible)

	if err != nil {
		log.Printf("[ERROR] scan failed, err:%v\n", err)
		return
	}
	userName, err := UserClient.GetNameById(ctx, announceItem.UserId)
	if err != nil {
		log.Printf("[ERROR] user.GetNameById failed, err:%v\n", err)
		return
	}
	timeLayout := "2022-01-01 00:01:02"
	createTime := time.Unix(announceItem.CreateTime, 0).Format(timeLayout)
	lastUpdateTime := time.Unix(announceItem.LastUpdateTime, 0).Format(timeLayout)
	resp = &announce.AnnounceResponse{
		AnnounceId:     announceItem.AnnounceId,
		Title:          announceItem.Title,
		UserName:       userName,
		Content:        announceItem.Content,
		CreateTime:     createTime,
		LastUpdateTime: lastUpdateTime,
	}
	log.Println("[INFO] return announce.GetAnnouncementById")
	return
}

// GetAllAnnouncements implements the ApiImpl interface.
func (s *ApiImpl) GetAllAnnouncements(ctx context.Context) (resp []*announce.AnnounceResponse, err error) {
	sqlStr := "select * from announcements where visible=1"
	rows, err := Db.Query(sqlStr)
	defer rows.Close()

	for rows.Next() {
		var item announcementFromDb
		err = rows.Scan(&item.AnnounceId, &item.Title, &item.UserId, &item.Content, &item.CreateTime, &item.LastUpdateTime, &item.visible)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		var userName string
		userName, err = UserClient.GetNameById(ctx, item.UserId)
		if err != nil {
			log.Printf("[ERROR] user.GetNameById failed, err:%v\n", err)
			return
		}
		timeLayout := "2022-01-01 00:01:02"
		createTime := time.Unix(item.CreateTime, 0).Format(timeLayout)
		lastUpdateTime := time.Unix(item.LastUpdateTime, 0).Format(timeLayout)
		resp = append(resp, &announce.AnnounceResponse{
			AnnounceId:     item.AnnounceId,
			Title:          item.Title,
			UserName:       userName,
			Content:        item.Content,
			CreateTime:     createTime,
			LastUpdateTime: lastUpdateTime,
		})
	}
	log.Println("[INFO] return announce.GetAllAnnouncements")

	return
}
