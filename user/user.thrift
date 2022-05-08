namespace go user

struct Identity {
	1: optional i32 userId
	2: optional string address
	3: optional i32 isAdmin
}

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
	string GetNameById(1: i32 userId, 2: Identity idt)	// 根据id查询用户名
	userInfo GetInfoById(1: i32 userId, 2: Identity idt)	// 根据id查询用户所有信息
	i32 InsertUser(1: insertUserInfo info, 2: Identity idt)	// 插入用户
	i32 UpdateUserInfo(1: i32 id, 2: insertUserInfo info,3: Identity idt)	// 更新用户信息
	i32 CheckUserName(1: string name)					//判断用户名是否被占用
	
}