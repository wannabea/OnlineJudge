namespace go user

struct userInfo {
	1: i32 userId
	2: string userName
	3: string passwd
	4: i32 createTime
	5: optional i32 lastLoginTime
	6: optional string realName
	7: optional string email
	8: optional string lastLoginIp 
	9: optional string signContent
	10: i32 isAdmin
}

struct insertUserInfo {
	1: string userName
	2: string passwd
	3: optional string realName
	4: optional string email
	5: optional string signContent
	6: i32 isAdmin
}

service User { 
	string GetNameById(1: i32 userId)
	userInfo GetInfoById(1: i32 userId)
	i32 InsertUser(1: insertUserInfo info)
	i32 updateUserInfo(1: insertUserInfo info)
	i32 checkUserName(1: string name)
}