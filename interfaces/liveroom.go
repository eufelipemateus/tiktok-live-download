package interfaces

type SigiState struct {
	LiveRoom struct {
		LoadingState struct {
			GetRecommendLive int `json:"getRecommendLive"`
			GetUserInfo      int `json:"getUserInfo"`
			GetUserStat      int `json:"getUserStat"`
		} `json:"loadingState"`
		NeedLogin          bool  `json:"needLogin"`
		ShowLiveGate       bool  `json:"showLiveGate"`
		IsAgeGateRoom      bool  `json:"isAgeGateRoom"`
		RecommendLiveRooms []any `json:"recommendLiveRooms"`
		LiveRoomStatus     int   `json:"liveRoomStatus"`
		LiveRoomUserInfo   struct {
			User struct {
				AvatarLarger string `json:"avatarLarger"`
				AvatarMedium string `json:"avatarMedium"`
				AvatarThumb  string `json:"avatarThumb"`
				ID           string `json:"id"`
				Nickname     string `json:"nickname"`
				SecUID       string `json:"secUid"`
				Secret       bool   `json:"secret"`
				UniqueID     string `json:"uniqueId"`
				Verified     bool   `json:"verified"`
				RoomID       string `json:"roomId"`
				Signature    string `json:"signature"`
				Status       int    `json:"status"`
				FollowStatus int    `json:"followStatus"`
			} `json:"user"`
			Stats struct {
				FollowingCount int `json:"followingCount"`
				FollowerCount  int `json:"followerCount"`
			} `json:"stats"`
			LiveRoom struct {
				CoverURL       string `json:"coverUrl"`
				SquareCoverImg string `json:"squareCoverImg"`
				Title          string `json:"title"`
				StartTime      int    `json:"startTime"`
				Status         int    `json:"status"`
				PaidEvent      struct {
					EventID  int `json:"event_id"`
					PaidType int `json:"paid_type"`
				} `json:"paidEvent"`
				LiveSubOnly   int `json:"liveSubOnly"`
				LiveRoomMode  int `json:"liveRoomMode"`
				GameTagID     int `json:"gameTagId"`
				LiveRoomStats struct {
					UserCount int `json:"userCount"`
				} `json:"liveRoomStats"`
				StreamData struct {
					PullData struct {
						Options struct {
							DefaultQuality struct {
								IconType   int    `json:"icon_type"`
								Level      int    `json:"level"`
								Name       string `json:"name"`
								Resolution string `json:"resolution"`
								SdkKey     string `json:"sdk_key"`
								VCodec     string `json:"v_codec"`
							} `json:"default_quality"`
							Qualities []struct {
								IconType   int    `json:"icon_type"`
								Level      int    `json:"level"`
								Name       string `json:"name"`
								Resolution string `json:"resolution"`
								SdkKey     string `json:"sdk_key"`
								VCodec     string `json:"v_codec"`
							} `json:"qualities"`
							ShowQualityButton bool `json:"show_quality_button"`
						} `json:"options"`
						StreamData string `json:"stream_data"`
					} `json:"pull_data"`
				} `json:"streamData"`
				StreamID          string `json:"streamId"`
				MultiStreamScene  int    `json:"multiStreamScene"`
				MultiStreamSource int    `json:"multiStreamSource"`
				HevcStreamData    struct {
					PullData struct {
						Options struct {
							DefaultQuality struct {
								IconType   int    `json:"icon_type"`
								Level      int    `json:"level"`
								Name       string `json:"name"`
								Resolution string `json:"resolution"`
								SdkKey     string `json:"sdk_key"`
								VCodec     string `json:"v_codec"`
							} `json:"default_quality"`
							Qualities []struct {
								IconType   int    `json:"icon_type"`
								Level      int    `json:"level"`
								Name       string `json:"name"`
								Resolution string `json:"resolution"`
								SdkKey     string `json:"sdk_key"`
								VCodec     string `json:"v_codec"`
							} `json:"qualities"`
							ShowQualityButton bool `json:"show_quality_button"`
						} `json:"options"`
						StreamData string `json:"stream_data"`
					} `json:"pull_data"`
				} `json:"hevcStreamData"`
			} `json:"liveRoom"`
		} `json:"liveRoomUserInfo"`
	} `json:"LiveRoom"`
	CurrentRoom struct {
		LoadingState struct {
			EnterRoom int `json:"enterRoom"`
		} `json:"loadingState"`
		RoomInfo          any    `json:"roomInfo"`
		AnchorID          string `json:"anchorId"`
		SecAnchorID       string `json:"secAnchorId"`
		AnchorUniqueID    string `json:"anchorUniqueId"`
		RoomID            string `json:"roomId"`
		HotLiveRoomInfo   any    `json:"hotLiveRoomInfo"`
		LiveType          string `json:"liveType"`
		ReportLinkType    string `json:"reportLinkType"`
		EnterRoomWithSSR  bool   `json:"enterRoomWithSSR"`
		PlayMode          string `json:"playMode"`
		IsGuestConnection bool   `json:"isGuestConnection"`
		IsMultiGuestRoom  bool   `json:"isMultiGuestRoom"`
		ShowLiveChat      bool   `json:"showLiveChat"`
		EnableChat        bool   `json:"enableChat"`
		IsAnswerRoom      bool   `json:"isAnswerRoom"`
		IsGateRoom        bool   `json:"isGateRoom"`
		RequestID         string `json:"requestId"`
		NtpDiff           int    `json:"ntpDiff"`
		FollowStatusMap   struct {
		} `json:"followStatusMap"`
	} `json:"CurrentRoom"`
}