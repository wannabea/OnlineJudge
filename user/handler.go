package main

import (
	"context"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/wannabea/OnlineJudge/user/db"
	"github.com/wannabea/OnlineJudge/user/kitex_gen/user"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// GetNameById implements the UserImpl interface.
func (s *UserImpl) GetNameById(ctx context.Context, userId int32) (resp string, err error) {
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
func (s *UserImpl) GetInfoById(ctx context.Context, userId int32) (resp *user.UserInfo, err error) {

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
func (s *UserImpl) InsertUser(ctx context.Context, info *user.InsertUserInfo) (resp int32, err error) {
	resp = 0
	sqlStr := "INSERT INTO users SET user_name=?, passwd=?, email=?,create_time=?,last_login_time=? ,is_admin=0 "
	timeUnix := time.Now().Unix()
	ret, err := Db.Exec(sqlStr, info.UserName, info.Passwd, info.Email, timeUnix, timeUnix)
	log.Printf("%v", info)
	if err != nil {
		log.Printf("insert failed, err:%v\n", err)
		return
	}
	resp = 1
	theID, err := ret.LastInsertId()
	if err != nil {
		log.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	log.Printf("insert success, the id is %d.\n", theID)
	return
}

// UpdateUserInfo implements the UserImpl interface.
func (s *UserImpl) UpdateUserInfo(ctx context.Context, id int32, info *user.UserInfo) (resp int32, err error) {
	resp = 0
	sqlStr := "UPDATE users SET user_name=?, passwd=?, email=?,last_login_time=? ,last_login_ip=?, sign_content=?, real_name=?, is_admin=? WHERE user_id=?"
	timeUnix := time.Now().Unix()
	ret, err := Db.Exec(sqlStr, info.UserName, info.Passwd, info.Email, timeUnix, info.LastLoginIp, info.SignContent, info.RealName, info.IsAdmin, id)
	log.Printf("%v", info)
	if err != nil {
		log.Printf("update failed, err:%v\n", err)
		return
	}
	resp = 1
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		log.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	log.Printf("update success, affected rows:%d\n", n)
	return
}

// CheckUserName implements the UserImpl interface.
func (s *UserImpl) CheckUserName(ctx context.Context, name string) (resp int32, err error) {
	// TODO: Your code here...
	return
}

// GetUserIdByUserName implements the UserImpl interface.
func (s *UserImpl) GetUserIdByUserName(ctx context.Context, userName string) (resp int32, err error) {
	sqlStr := "select user_id from users where user_name=?"
	var userId int32
	err = Db.QueryRow(sqlStr, userName).Scan(&userId)
	log.Println(userId)
	if err != nil {
		log.Printf("scan failed, err:%v\n", err)
		return -1, err
	}
	resp = userId
	return
}
