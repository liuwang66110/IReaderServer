package friend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"controllers"
	"models"
	"strconv"
	"fmt"
)

func FriendAdd(c *gin.Context)  {
	tmp, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "请先登陆", nil))
		return
	}
	if c.PostForm("friend_id") == "" {
		c.JSON(http.StatusBadRequest, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "缺少参数:friend_id",nil))
		return
	}
	friendID, idError := strconv.ParseUint(c.PostForm("friend_id"), 10, 64)
	if (idError != nil) {
		c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "id错误", nil))
		return
	}
	user := tmp.(models.User)

	friend := models.UserFriend{
		UserID: friendID,
		FriendID: user.ID,
		Status: models.FRIEND_ST_WATTING,
	}

	db := models.Instance()
	var count int
	if err := db.Model(models.UserFriend{}).Where(models.UserFriend{UserID: friendID, FriendID: user.ID}).Count(&count).Error; err != nil {
		c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_SERVER_ERROR, nil))
		return
	}
	if (count > 0) {
		c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "已申请", nil))
		return
	}
	if err := db.Create(&friend).Error; err != nil {
		c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_SERVER_ERROR, nil))
		return
	}
	c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_SUCCESS, "添加成功", nil))
	return
}
