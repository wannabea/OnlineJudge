namespace go announcement

struct AnnounceRequest {
  1: i32 announce_id
}

struct AnnounceResponse {
  1: i32 announceId
  2: string title
  3: string userName
  4: string content
  5: string create_time
  6: string last_update_time
}

service Api {
    AnnounceResponse GetAnnouncementById (1: AnnounceRequest req)
	list<AnnounceRequest> GetAllAnnouncements ()
}

