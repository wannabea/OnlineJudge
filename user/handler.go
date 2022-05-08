package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/wannabea/OnlineJudge/user/db"
	"github.com/wannabea/OnlineJudge/user/kitex_gen/user"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// GetNameById implements the UserImpl interface.
func (s *UserImpl) GetNameById(ctx context.Context, userId int32, idt *user.Identity) (resp string, err error) {
	sqlStr := "select user_name from users where user_id=?"
	err = Db.QueryRow(sqlStr, userId).Scan(&resp)
	if err != nil {
		log.Printf("scan failed, err:%v\n", err)
		return
	}
	log.Println("[INFO] return from user.GetNameById")
	return
}

func checkAccess(id int32, idt user.Identity) int {
	access := 0
	if idt.UserId == id {
		access = 1
	} else if idt.IsAdmin == int32(1) {
		access = 1
	}
	return access
}

// GetInfoById implements the UserImpl interface.
func (s *UserImpl) GetInfoById(ctx context.Context, userId int32, idt *user.Identity) (resp *user.UserInfo, err error) {
	// if checkAccess(userId, *idt) == 0 {
	// 	log.Println("[ERROR] have no access")
	// 	return
	// }
	sqlStr := "select * from users where user_id=?"
	var item user.UserInfo
	// err = Db.QueryRow(sqlStr, userId).Scan(&resp.UserId, &resp.UserName, &resp.Passwd, &resp.CreateTime, &resp.LastLoginTime, &resp.RealName, &resp.Email, &resp.LastLoginIp, &resp.SignContent, &resp.IsAdmin)
	err = Db.QueryRow(sqlStr, userId).Scan(&item.UserId, &item.UserName, &item.Passwd, &item.CreateTime, &item.LastLoginTime, &item.RealName, &item.Email, &item.LastLoginIp, &item.SignContent, &item.IsAdmin)
	resp = &item
	if err != nil {
		log.Printf("scan failed, err:%v\n", err)
		return
	}
	log.Println("[INFO] return from user.GetInfoById")
	return
}

// InsertUser implements the UserImpl interface.
func (s *UserImpl) InsertUser(ctx context.Context, info *user.InsertUserInfo, idt *user.Identity) (resp int32, err error) {
	
	return
}

// UpdateUserInfo implements the UserImpl interface.
func (s *UserImpl) UpdateUserInfo(ctx context.Context, id int32, info *user.InsertUserInfo, idt *user.Identity) (resp int32, err error) {
	// TODO: Your code here...
	return
}

// CheckUserName implements the UserImpl interface.
func (s *UserImpl) CheckUserName(ctx context.Context, name string) (resp int32, err error) {
	// TODO: Your code here...
	return
}
