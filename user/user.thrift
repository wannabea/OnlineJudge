namespace go user

struct Identity {
	1: i32 userId
	3: i32 isAdmin
}

struct userInfo {
	1: i32 userId
	2: string userName
	3: string passwd
	4: i32 createTime
	5: i32 lastLoginTime
	6: string realName
	7: string email
	8: string lastLoginIp 
	9: string signContent
	10: i32 isAdmin
}

struct insertUserInfo {
	1: string userName
	2: string passwd
	3: string email
}

service User { 
	string GetNameById(1: i32 userId)	// 根据id查询用户名
	userInfo GetInfoById(1: i32 userId)	// 根据id查询用户所有信息
	i32 GetUserIdByUserName(1: string userName)
	i32 InsertUser(1: insertUserInfo info)	// 插入用户
	i32 UpdateUserInfo(1: i32 id, 2: userInfo info)	// 更新用户信息
	i32 CheckUserName(1: string name)					//判断用户名是否被占用
}