package mysql

const (
	USER_INFO_TABLE = "info"
)

type UserInfo struct {
	Uid  int64  `json:"uid"`
	Data []byte `json:"data"`
	Timestamp
}

type UserInfoRepo struct{}

func (userInfoRepo *UserInfoRepo) GetUserInfo(uid int64) (*UserInfo, error) {
	return nil, nil
}
