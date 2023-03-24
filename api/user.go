package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"gomock/service"
	"net/http"
	"strconv"
)

type NewUserBasicInfoReq struct {
	Uid           int64 `json:"uid"`
	CurStage      int32 `json:"cur_stage"`
	MaincityLevel int32 `json:"maincity_level"`
}

func GetUserBasicInfo(c *gin.Context) {
	uidStr := c.Query("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)

	if err != nil {
		c.JSON(http.StatusOK, CodeParamErr.WithErr(err))
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, CodeInternalServiceErr.WithErr(err))
		return
	}

	re, err := service.GetUserBasicInfo(uid)

	if err != nil {
		c.JSON(http.StatusOK, CodeInternalServiceErr.WithErr(err))
		return
	}

	c.JSON(
		http.StatusOK,
		GetUserBasicInfoResp{
			service.UserBasicInfo{
				Uid:           uid,
				CurStage:      re.CurStage,
				MaincityLevel: re.MaincityLevel,
			},
		})
}

func GetUserInfo(c *gin.Context) {
	uid := c.Query("uid")
	lang := c.DefaultQuery("language", "en")

	uidInt64, _ := strconv.ParseInt(uid, 10, 64)

	re, err := service.GetUserInfo(uidInt64, lang)
	if err != nil {
		c.JSON(http.StatusOK, CodeInternalServiceErr.WithErr(err))
	}

	c.JSON(http.StatusOK, GetUserInfoResp{
		service.UserInfo{
			Uid:            re.Uid,
			UserName:       re.UserName,
			Avatar:         re.Avatar,
			Kingdom:        service.Kingdom{},
			Guild:          service.Guild{},
			BadgeUrl:       "",
			SubTitleList:   nil,
			AvatarFrameUrl: "",
			BubbleConfigs:  service.BubbleConfigs{},
			VipLevel:       0,
			ShowVip:        false,
			Level:          re.Level,
			GameExtra:      "",
			Sex:            re.Sex,
			EmblemUrls:     re.EmblemUrls,
			CreateTime:     re.CreateTime,
			TextColor:      re.TextColor,
			UserType:       re.UserType,
		},
	})
}

func GetUserBatchInfo(c *gin.Context) {}
func GetUserGroups(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
func GetUserFriends(c *gin.Context)    {}
func GetUserBlackList(c *gin.Context)  {}
func GetUserFriendReqs(c *gin.Context) {}

//-- POSTS -- //

func PostUserCreate(c *gin.Context) {
	var newUserBasicInfoReq NewUserBasicInfoReq
	err := json.NewDecoder(c.Request.Body).Decode(&newUserBasicInfoReq)

	if err != nil {
		c.JSON(http.StatusOK, CodeParamErr.WithErr(err))
		return
	}

	// service
	err = service.NewUserBasicInfo(newUserBasicInfoReq.Uid, newUserBasicInfoReq.CurStage, newUserBasicInfoReq.MaincityLevel)

	if err != nil {
		c.JSON(http.StatusOK, CodeDBErr.WithErr(err))
		return
	}

	c.JSON(http.StatusOK, CodeOk)
}
func PostUserAddFriend(c *gin.Context)    {}
func PostUserAcceptFriend(c *gin.Context) {}
func PostUserDelFriend(c *gin.Context)    {}
func PostUserAddBlackList(c *gin.Context) {}
func PostUserDelBlackList(c *gin.Context) {}
