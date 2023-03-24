package service

type UserBasicInfo struct {
	Uid           int64 `json:"uid" db:"uid"`
	CurStage      int32 `json:"cur_stage" db:"cur_stage"`
	MaincityLevel int32 `json:"maincity_level" db:"maincity_level"`
}

type Kingdom struct {
	StoryId   int64  `json:"story_id"`
	KingdomId int64  `json:"kingdom_id"`
	AvatarUrl string `json:"avatar_url"`
}
type Guild struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	AbbrName  string `json:"abbr_name"`
}

type SubTitleItem struct {
	Key     string
	Content string
	BgUrl   string
}
type BubbleConfig struct {
	Url        string
	EdgeInsets []int32
}

type BubbleConfigs struct {
	LeftNormal   BubbleConfig
	leftPressed  BubbleConfig
	RightNormal  BubbleConfig
	RightPressed BubbleConfig
}

type UserInfo struct {
	Uid            int64          `json:"uid"`
	UserName       string         `json:"user_name"`
	Avatar         string         `json:"avatar"`
	Kingdom        Kingdom        `json:"kingdom"`
	Guild          Guild          `json:"guild"`
	BadgeUrl       string         `json:"badge_url"`
	SubTitleList   []SubTitleItem `json:"sub_title_list"`
	AvatarFrameUrl string         `json:"avatar_frame_url"`
	BubbleConfigs  BubbleConfigs  `json:"bubble_configs"`
	VipLevel       int32          `json:"vip_level"`
	ShowVip        bool           `json:"show_vip"`
	Level          int32          `json:"level"`
	GameExtra      string         `json:"game_extra"`
	Sex            int32          `json:"sex"`
	EmblemUrls     []string       `json:"emblem_urls"`
	CreateTime     int64          `json:"create_time"`
	TextColor      string         `json:"text_color"`
	UserType       int32          `json:"user_type"`
}

type GroupInfo struct {
	AtAllPerDay  int32   `json:"at_all_per_day"`
	GroupAvatar  string  `json:"group_avatar"`
	GroupId      int64   `json:"group_id"`
	GroupName    string  `json:"group_name"`
	GroupSubType int32   `json:"group_sub_type"`
	GroupType    int32   `json:"group_type"`
	ManagerList  []int64 `json:"manager_list"`
	ServerId     string  `json:"server_id"`
}
