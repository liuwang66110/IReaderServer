package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"controllers"
	"strings"
	"time"
	"models"
	"utils"
)

func UserRegister(c *gin.Context)  {
	keys := []string{"name", "password"}

	for _, v := range keys {
		if c.PostForm(v) == "" {
			c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "缺少参数:" + v,nil))
			return
		}
	}

	user := models.User{
		Password:    utils.GenPwd(c.PostForm("password")),
		Token:       utils.RandMd5(),
		Name:        c.PostForm("name"),
		ExpiredAt:   time.Now().Add(models.USER_TOKEN_DURATION),
		Status:      models.USER_ST_ENABLE,
		CreatedAt:   models.StdTime(time.Now()),
		LastLoginAt: models.StdTime(time.Now()),
	}
	db := models.Instance()
	tx := db.Begin()
	err := tx.Create(&user).Error
	if err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		c.JSON(controllers.Rsp(controllers.OK_INSERT_FAILED, nil, "用户名重复"))
		tx.Rollback()
		return
	}
	if err != nil {
		utils.LogInstance().Println(err.Error())
		c.JSON(http.StatusInternalServerError, controllers.SetRsp(controllers.OK_SERVER_ERROR, nil))
		tx.Rollback()
		return
	}

	userInfo := models.UserInfo{
		UserID: user.ID,
		Name: utils.GetRandomString(12),
	}
	if err := tx.Create(&userInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, controllers.SetRsp(controllers.OK_SERVER_ERROR, nil))
		tx.Rollback()
		return
	}
	tx.Commit()
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["name"] = userInfo.Name
	data["token"] = user.Token
	c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_INSERT_SUCCESS, data))
	return
}