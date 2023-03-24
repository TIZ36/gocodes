package mysql

import "time"

type Timestamp struct {
	CreateTime time.Time
	UpdateTime time.Time
}
type DbUserInfo struct {
	Uid  int64
	Data []byte

	Timestamp
}

type DbUserGroupMap struct {
	Id  int64
	Gid int64
	Uid int64

	Timestamp
}

type DbUserFriend struct {
	OneUid     int64
	AnoUid     int64
	RelateHash string

	Timestamp
}

type DbUserBlacklist struct {
	Id         int64
	OneUid     int64
	AnoUid     int64
	RelateHash string
	BlackState int32

	Timestamp
}

type DbGroupOwner struct {
	Gid      int64
	OwnerUid int64

	Timestamp
}

type DbGroupConfig struct {
	GroupId      int64
	PositionData []byte

	Timestamp
}

type DbGroupInfo struct {
	GroupId      int64
	GroupName    int64
	GroupAvatar  string
	GroupType    int32
	GroupSubType int32
	ServerId     int32
	ManagerList  []byte
	AtAllPerDay  int32

	Timestamp
}
