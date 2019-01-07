package friend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"controllers"
	"models"
	"strconv"
)

func FriendList(c *gin.Context) {
	tmp, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "请先登陆", nil))
		return
	}

	page, _ := strconv.Atoi(c.PostForm("page"))
	page_size, _ := strconv.Atoi(c.PostForm("pageSize"))
	if page < -1 {
		c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "页数错误", nil))
		return
	}

	user := tmp.(models.User)

	db := models.Instance()
	var count int
	db = db.Model(models.UserFriend{}).Where(models.UserFriend{UserID: user.ID, Status: models.FRIEND_ST_AGREE}).Count(&count)

	if page > -1 {
		db = db.Offset(page * page_size).Limit(page_size)
	}

	var friendList []models.UserFriend
	if err := db.Find(&friendList).Error; err != nil {
		c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_SERVER_ERROR, nil))
		return
	}
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["list"] = friendList
	var totalPage int
	if page == 0 || page_size == 0 {
		data["page"] = 0
	} else {
		if count % page_size == 0 {
			totalPage = count / page_size
		} else {
			totalPage = count / page_size + 1
		}
		data["page"] = totalPage
	}
	data["page"] = totalPage
	c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_INSERT_SUCCESS, data))
	return
}
