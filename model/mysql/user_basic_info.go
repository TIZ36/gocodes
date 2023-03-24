package mysql

import (
	"fmt"
	"gomock/resource"
)

const (
	TABLE = "basic_info"
)

type BasicInfo struct {
	Uid           int64
	CurStage      int32
	MaincityLevel int32
	Timestamp
}

type UserBasicInfoRepo struct{}

// FindBasicInfoByUid 通过uid查询用户基本信息 /
func (basicInfoRepo *UserBasicInfoRepo) FindBasicInfoByUid(uid int64) (*BasicInfo, error) {
	var basicInfo BasicInfo
	sql := fmt.Sprintf("SELECT * FROM %s WHERE `uid` = %v", TABLE, uid)
	err := resource.AppCtx.MysqlClient.QueryRow(sql).Scan(
		&basicInfo.Uid,
		&basicInfo.CurStage,
		&basicInfo.MaincityLevel,
		&basicInfo.CreateTime,
		&basicInfo.UpdateTime)

	if err != nil {
		return nil, err
	}
	return &basicInfo, nil
}

func (basicInfoRepo *UserBasicInfoRepo) InsertBasicInfo(dbBasicInfo BasicInfo) error {
	_ = fmt.Sprintf(
		"INSERT INTO %s (`uid`, `cur_stage`, `maincity_level`) VALUES (%v, %v, %v)",
		TABLE,
		dbBasicInfo.Uid,
		dbBasicInfo.CurStage,
		dbBasicInfo.MaincityLevel,
	)

	stm, err := resource.AppCtx.MysqlClient.Prepare(
		"INSERT INTO `basic_info` (`uid`, `cur_stage`, `maincity_level`) VALUES (?, ?, ?)")

	defer stm.Close()

	if err != nil {
		return err
	}

	_, err = stm.Exec(dbBasicInfo.Uid, dbBasicInfo.CurStage, dbBasicInfo.MaincityLevel)

	return err
}
