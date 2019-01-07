package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"controllers"
	"time"
	"models"
	"utils"
	"github.com/jinzhu/gorm"
	"fmt"
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
	var selectUser models.User
	if err := db.Model(models.User{}).Where(models.User{Name: c.PostForm("name")}).Last(&selectUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			insertUser(c, db, user)
			return
		}
		c.JSON(http.StatusInternalServerError, controllers.SetRsp(controllers.OK_SERVER_ERROR, nil))
		return
	}
	fmt.Println(selectUser)
	if selectUser.Status == models.USER_ST_DISABLE || selectUser.Status == models.USER_ST_DELETED {
		insertUser(c, db, user)
		return
	}
	c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "用户已存在", nil))
	return
}

func insertUser(c *gin.Context, db *gorm.DB, user models.User)  {
	tx := db.Begin()
	if err := tx.Create(&user).Error; err != nil {
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