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
	resp = &announce.AnnounceInfo{
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
func (s *ApiImpl) GetAllAnnouncements(ctx context.Context) (resp []*announce.AnnounceInfo, err error) {
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
		resp = append(resp, &announce.AnnounceInfo{
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

// InsertAnnouncement implements the ApiImpl interface.
func (s *ApiImpl) InsertAnnouncement(ctx context.Context, info *announce.UpdateAnouncementRequest, idt *announce.Identity) (resp int32, err error) {
	if *idt.IsAdmin != int32(1) || idt.UserId == nil {
		log.Println("[ERROR] have no access")
		return 0, nil
	}
	sqlStr := "INSERT INTO announcements SET (title=?, user_id=?, content=?, create_time=?, last_update_time=?, visible=?)"
	ret, err := Db.Exec(sqlStr, info.Title, info.UserId, info.Content, info.CreateTime, info.LastUpdateTime, info.Visible)
	if err != nil {
		log.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		log.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	log.Printf("insert success, the id is %d.\n", theID)
	return
}

// UpdateAnnouncement implements the ApiImpl interface.
func (s *ApiImpl) UpdateAnnouncement(ctx context.Context, announceId int32, info *announce.UpdateAnouncementRequest, idt *announce.Identity) (resp int32, err error) {
	if *idt.IsAdmin != int32(1) || idt.UserId == nil {
		log.Println("[ERROR] have no access")
		return 0, nil
	}
	sqlStr := "UPDATE announcements SET (title=?, user_id=?, content=?, create_time=?, last_update_time=?, visible=? where announce_id=?)"
	ret, err := Db.Exec(sqlStr, info.Title, info.UserId, info.Content, info.CreateTime, info.LastUpdateTime, info.Visible, announceId)
	if err != nil {
		log.Printf("update failed, err:%v\n", err)
		return 0, err
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		log.Printf("get RowsAffected failed, err:%v\n", err)
		return 0, err
	}
	log.Printf("update success, affected rows:%d\n", n)
	return 1, nil
}

// DeleteAnnouncement implements the ApiImpl interface.
func (s *ApiImpl) DeleteAnnouncement(ctx context.Context, id int32, idt *announce.Identity) (resp int32, err error) {
	if *idt.IsAdmin != int32(1) || idt.UserId == nil {
		log.Println("[ERROR] have no access")
		return 0, nil
	}
	sqlStr := "DELETE  FROM announcements where announce_id=?)"
	ret, err := Db.Exec(sqlStr, id)
	if err != nil {
		log.Printf("delete failed, err:%v\n", err)
		return 0, err
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		log.Printf("get RowsAffected failed, err:%v\n", err)
		return 0, err
	}
	log.Printf("delete success, affected rows:%d\n", n)

	return 1, nil
}

// HideAnnouncement implements the ApiImpl interface.
func (s *ApiImpl) HideAnnouncement(ctx context.Context, op int32, idt *announce.Identity) (resp int32, err error) {
	if *idt.IsAdmin != int32(1) || idt.UserId == nil {
		log.Println("[ERROR] have no access")
		return 0, nil
	}
	sqlStr := "UPDATE announcements SET (visible=?)"
	ret, err := Db.Exec(sqlStr, op)
	if err != nil {
		log.Printf("update failed, err:%v\n", err)
		return 0, err
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		log.Printf("get RowsAffected failed, err:%v\n", err)
		return 0, err
	}
	log.Printf("update success, affected rows:%d\n", n)
	return 1, nil
}
