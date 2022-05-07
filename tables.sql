// Announcements

CREATE TABLE announcements (	--公告表
	`announce_id` INT AUTO_INCREMENT,		-- 主键id
	`title` VARCHAR(50) NOT NULL,			-- 公告title
	`user_id` INT NOT NULL,					-- 发布用户id
	`content` varchar(10000),				-- 内容文本
	`create_time` INT NOT NULL,				-- 创建时间
	`last_update_time` INT NOT NULL			-- 最后修改时间
)

INSERT INTO announcements (title, user_id, content, create_time, last_update_time) VALUES(
	"测试公告1",1,"test",1651892698,1651892698
)

CREATE TABLE problems (			-- 题目表
	`problem_id` INT AUTO_INCREMENT,		-- 主键id
	`title` varchar(50) NOT NULL,			-- 问题title
	`user_id` INT NOT NULL,					-- 创建者id
	`is_visible` INT,						-- 是否可见
	`configs` VARCHAR(3000) NOT NULL,		-- 配置信息（题目的时间限制空间限制等等）json格式
	`tags` 	VARCHAR(300) NOT NULL,			-- 标签集合 json格式
	`source` varchar(300) NOT NULL,			-- 题目来源
	`create_time` INT NOT NULL,				-- 创建时间
	`last_update_time` INT NOT NULL			-- 最后更新时间
)

CREATE TABLE testcase (						-- 测试数据表
	`test_id` INT AUTO_INCREMENT,			-- 主键id
	`problem_id` INT NOT NULL,				-- 对应问题
	`in` varchar(300000) NOT NULL,			-- 输入文本
	`out` varchar(300000) NOT NULL			-- 输出文本
)

CREATE TABLE submits (						-- 提交信息
	`submits_id` BIGINT AUTO_INCREMENT,		-- 主键id
	`problem_id` INT NOT NULL,				-- 问题id
	`submit_time` INT NOT NULL,				-- 提交时间
	`status` INT NOT NULL,					-- 返回状态	int映射
	`consume_time` INT,						-- 消耗时间
	`consume_memory` INT,					-- 消耗空间
	`language` varchar(30) NOT NULL,		-- 使用语言
	`user_id` INT NOT NULL,					-- 提交用户
	`code` varchar(300000) NOT NULL,		-- 提交代码
	`score` DOUBLE NOT NULL					-- 通过数据得分
	`contest_id` INT,						-- 属于比赛id
)

CREATE TABLE contests (						-- 比赛表
	`contest_id` INT AUTO_INCREMENT,		-- 主键id
	`create_time` INT NOT NULL,				-- 创建时间	
	`last_update_time` INT NOT NULL,		-- 最后更新时间
	`is_lock` INT NOT NULL,					-- 是否加锁
	`passewd` varchar(30),					-- 密码
	`hints` varchar(300000),				-- 提示信息
	`problems` varchar(3000),				-- 问题列表 json集合
	`ranking` varchar(300000)				-- 比赛排行榜 json列表
)

CREATE TABLE users (						-- 用户表
	`user_id` INT AUTO_INCREMENT,			-- 用户id
	`user_name` INT PRIMARY KEY,			-- 用户名
	`passwd`	varchar(100) NOT NULL,		-- 密码 MD5
	`create_time` INT NOT NULL,				-- 注册时间
	`last_login_time` INT NOT NULL,			-- 最后登陆时间
	`real_name` varchar(30),				-- 真实姓名
	`email` varchar(30) NOT NULL,			-- email
	`last_login_ip` NOT NULL,				-- 最后登录ip
	`sign_content` varchar(100) 			-- 签名
)



CREATE TABLE announcements (
	`announce_id` INT PRIMARY KEY AUTO_INCREMENT,		
	`title` VARCHAR(50) NOT NULL,		
	`user_id` INT NOT NULL,				
	`content` varchar(10000),			
	`create_time` INT NOT NULL,			
	`last_update_time` INT NOT NULL		
)
