package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouter(c *gin.Context) {
	t := c.Query("type")

	switch t {
	// -- users -- /
	case "user.basicinfo":
		GetUserBasicInfo(c)
	case "user.info":
		GetUserInfo(c)
	case "user.batchInfo":
		GetUserBatchInfo(c)
	case "user.groups":
		GetUserGroups(c)
	case "user.friends":
		GetUserFriends(c)
	case "user.blacklist":
		GetUserBlackList(c)
	case "user.friendReqs":
		GetUserFriendReqs(c)
		// -- groups -- /
	case "group.info":
		GetGroupInfo(c)
	case "group.config":
		GetGroupConfig(c)
	default:
		c.JSON(http.StatusNotFound, CodeParamErr)
		return
	}
}
