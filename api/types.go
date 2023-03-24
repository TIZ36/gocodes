package api

import (
	"go/types"
	"gomock/service"
)

type CommonApiResp[T any] struct {
	ErrCode
	Data T `json:"data"`
}

type UserIn struct {
	Lang string `json:"language"`
	Uid  string `json:"uid"`
}

type UserOut struct {
	Avatar string `json:"avatar"`
	Id     int64  `json:"id"`
	Name   string `json:"name"`
}

type GroupUser struct {
	Avatar      string `json:"avatar"`
	GuildName   string `json:"guild_name"`
	KingdomIcon string `json:"kindom_icon"`
	Uid         int64  `json:"uid"`
	UserName    string `json:"user_name"`
}

type Group struct {
	Avatar string `json:"avatar"`
	Id     int64  `json:"id"`
	Lang   string `json:"language"`
}

// ----- end of common part -----//

// GET
// path: /chat?type=user.basicinfo&uid=35878897715245056&language=en
type (
	GetUserBasicInfoReq struct {
		Uid      int64  `query:"uid"`
		Language string `query:"language, optional"`
	}
	GetUserBasicInfoResp struct {
		service.UserBasicInfo
	}
)

// GET
// path: /chat?type=user.info&uid=362367283101655043&language=en
type (
	GetUserInfoReq struct {
		Uid  int64  `query:"uid"`
		Lang string `query:"language"`
	}

	GetUserInfoResp struct {
		service.UserInfo
	}
)

// GET
// /chat?type=user.batchInfo&users=%5B%7B%22language%22%3A%22en%22%2C%22uid%22%3A35878897715245056%7D%2C%7B
// %22language%22%3A%22en%22%2C%22uid%22%3A36237131768655872%7D%5D
type (
	BatchGetUserInfoReq struct {
		Users string `query:"users"` // www_url_encoded
	}
	BatchGetUserInfoResp []service.UserInfo
)

// POST
// path: /chat?type=user.create
type (
	UserCreateReturn struct {
		UserBasicInfo service.UserBasicInfo `json:"basicinfo"`
		UserInfo      service.UserInfo      `json:"info"`
	}
	UserCreateReq struct {
		Uid int64 `json:"uid,optional"`
	}
	UserCreateResp struct {
		CommonApiResp[UserCreateReturn]
	}
)

// POST
// path: /chat?type=user.add_friend
type (
	UserAddFriendReq struct {
		FromUid int64 `json:"from_uid"`
		ToUid   int64 `json:"to_uid"`
		Msg     int64 `json:"msg"`
	}
	UserAddFriendResp struct {
		CommonApiResp[map[string]string]
	}
)

// path: /chat?type=user.friend_reqs&uid=362367283101655043
type (
	GetUserFriendReqsReq struct {
		Uid int64 `query:"uid"`
	}
	GetUserFriendReqsResp []string
)

// path: /chat?type=user.accept_friend
type (
	UserAcceptFriendReq struct {
		Uid       int64  `json:"uid"`
		RequestId string `json:"request_id"`
	}
	UserAcceptFriendResp struct {
		CommonApiResp[types.Nil]
	}
)

// POST
// path: /chat?type=user.accept_friend
type (
	UserDelFriendReq struct {
		Uid int64 `json:"uid"`
		Fid int64 `json:"fid"`
	}
	UserDelFriendResp struct {
		CommonApiResp[types.Nil]
	}
)

// GET
// path: /chat?type=user.friends&uid=362367283101655043&language=en
type (
	GetUserFriendReq struct {
		Uid      int64  `query:"uid"`
		Language string `query:"language"`
	}
	GetUserFriendResp []UserOut
)

// POST
// path: /chat?type=user.add_blacklist
type (
	UserAddBlacklistReq struct {
		Uid      int64 `json:"uid"`
		BlackUid int64 `json:"black_uid"`
	}
	UserAddBlacklistResp struct {
		CommonApiResp[types.Nil]
	}
)

// POST
// path: /chat?type=user.del_blacklist
type (
	UserDelBlacklistReq struct {
		Uid      int64 `json:"uid"`
		BlackUid int64 `json:"black_uid"`
	}
	UserDelBlacklistResp struct {
		CommonApiResp[types.Nil]
	}
)

// GET
// path: /chat?type=user.blacklist&uid=35878897715245056
type (
	UserBlacklistReq struct {
		Uid int64 `query:"uid"`
	}
	UserBlacklistResp struct {
		Blacklist   []int64 `json:"blacklist"`
		Blacklisted []int64 `json:"blacklisted"`
	}
)

// POST
// path: /chat?type=group.create
type (
	GroupCreateReq  struct{}
	GroupCreateResp struct{}
)

// GET
// path: /chat?type=group.info&groupId=43072057616887808&language=en
type (
	GetGroupInfoReq struct {
		GroupId  int64  `query:"group_id"`
		Language string `query:"language"`
	}
	GetGroupInfoResp struct {
		service.GroupInfo
	}
)

// GET
// path: /chat?type=group.list&page=0&page_size=2
type (
	GetGroupListReq struct {
		Page     int32 `json:"page"`
		PageSize int32 `json:"page_size"`
	}
	GetGroupListResp []service.GroupInfo
)

// POST
// path: /chat?type=group.join
type (
	GroupJoinReq struct {
		GroupId int64 `json:"group_id"`
		Uid     int64 `json:"uid"`
	}
	GroupJoinResp struct {
		CommonApiResp[types.Nil]
	}
)

// POST
// path: /chat?type=group.leave
type (
	GroupLeaveReq struct {
		GroupId int64 `json:"group_id"`
		Uid     int64 `json:"uid"`
	}
	GroupLeaveResp struct {
		CommonApiResp[types.Nil]
	}
)

// GET
// path: /chat?type=user.groups&uid=35878897715245056&language=en
type (
	GetUserGroupsReq struct {
		Uid      int64  `query:"uid"`
		Language string `query:"language"`
	}
	GetUserGroupsResp []service.GroupInfo
)

// GET
// path: /chat?type=group.users&groupId=36242349231173632&language=en
type (
	GetGroupUsersReq struct {
		GroupId  int64  `json:"group_id"`
		Language string `json:"language"`
	}
	GetGroupUsersResp []GroupUser
)
