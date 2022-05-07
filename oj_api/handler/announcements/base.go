package announcements

type Announcement struct {
	Title          string `thrift:"title,1" json:"title"`
	UserName       string `thrift:"userName,2" json:"userName"`
	Content        string `thrift:"content,3" json:"content"`
	CreateTime     string `thrift:"create_time,4" json:"create_time"`
	LastUpdateTime string `thrift:"last_update_time,5" json:"last_update_time"`
}

type AnnouncementSimple struct {
	Title          string `thrift:"title,1" json:"title"`
	UserName       string `thrift:"userName,2" json:"userName"`
	CreateTime     string `thrift:"create_time,4" json:"create_time"`
	LastUpdateTime string `thrift:"last_update_time,5" json:"last_update_time"`
}

// type AnnouncementList struct {
// 	announcements []AnnouncementSimple
// }