namespace go announce

struct Identity {
	1: optional i32 userId
	2: optional string address
	3: optional i32 isAdmin
}

struct AnnounceRequest {
  1: i32 announce_id
}

struct AnnounceInfo {
  1: i32 announceId
  2: string title
  3: string userName
  4: string content
  5: string create_time
  6: string last_update_time
}

service Api {
    AnnounceInfo GetAnnouncementById (1: AnnounceRequest req)
	list<AnnounceInfo> GetAllAnnouncements ()
	i32 InsertAnnouncement(1: AnnounceInfo, 2: Identity)
	i32 UpdateAnnouncement(1: AnnounceInfo, 2: Identity)
	i32 DeleteAnnouncement(1: AnnounceInfo, 2: Identity)
	i32 HideAnnouncement(1: AnnounceInfo, 2: Identity)
}

