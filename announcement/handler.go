package main

import (
	"context"
	"database/sql"
	"strconv"
	"fmt"
	
	"github.com/OnlineJudge/announcement/kitex_gen/announcement"
	_ "github.com/go-sql-driver/mysql"
)

// ApiImpl implements the last service interface defined in the IDL.
type ApiImpl struct{}
var db *sql.DB

// GetAnnouncementById implements the ApiImpl interface.

type announcementFromDb struct {
	AnnounceId    	int
	Title          	string 
	UserId       	string 
	Content        	string 
	CreateTime     	int 
	LastUpdateTime 	int 
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
		Title: announce.Title,
		UserName: strconv.Itoa(announce.UserId),
		Content: announce.Content,

	}
	return
}

// GetAllAnnouncements implements the ApiImpl interface.
func (s *ApiImpl) GetAllAnnouncements(ctx context.Context) (resp []*announcement.AnnounceRequest, err error) {
	// TODO: Your code here...
	return
}
