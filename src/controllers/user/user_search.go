package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"controllers"
	"models"
	"strings"
)

func UserSearch(c *gin.Context) {
	tmp, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "请先登陆", nil))
		return
	}

	if (c.PostForm("searchKey") == "") {
		c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "请输入搜索内容", nil))
		return
	}

	user := tmp.(models.User)

	searchKey := c.PostForm("searchKey")
	var userList []models.UserInfo

	db := models.Instance()
	if err := db.Model(models.UserInfo{}).Where("name LIKE ?", "%" + searchKey + "%").Not(models.UserInfo{UserID: user.ID}).
		Find(&userList).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "没有该用户", nil))
			return
		}
		c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_SERVER_ERROR, nil))
		return
	}
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["list"] = userList
	c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_INSERT_SUCCESS, data))
	return
}
