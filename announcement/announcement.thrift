namespace go announcement

struct AnnounceRequest {
  1: i32 announce_id
}

struct AnnounceResponse {
  1: string title
  2: string userName
  3: string content
  4: string create_time
  5: string last_update_time
  6: i32 announce_id
}

service Api {
    AnnounceResponse GetAnnouncementById (1: AnnounceRequest req)
	list<AnnounceRequest> GetAllAnnouncements ()
}

