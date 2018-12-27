package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"controllers"
	"strings"
	"time"
	"strconv"
	"models"
	"utils"
)

func UserRegister(c *gin.Context)  {
	keys := []string{"name", "password"}

	for _, v := range keys {
		if c.PostForm(v) == "" {
			c.JSON(http.StatusBadRequest, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "缺少参数:" + v,nil))
			return
		}
	}

	var admin models.User
	var mobile uint64
	if (c.PostForm("mobile") != "") {
		mobile, _ = strconv.ParseUint(c.PostForm("mobile"), 10, 64)
	}

	admin = models.User{
		Password:    utils.GenPwd(c.PostForm("password")),
		Token:       utils.RandMd5(),
		Name:        c.PostForm("name"),
		Mobile:      mobile,
		ExpiredAt:   time.Now().Add(models.USER_TOKEN_DURATION),
		Status:      models.USER_ST_ENABLE,
		CreatedAt:   time.Now(),
		LastLoginAt: models.StdTime(time.Now()),
	}
	db := models.Instance()
	// err := db.Where(models.CorpAdmin{Name: c.PostForm("name")}).FirstOrCreate(&admin).Error
	err := db.Create(&admin).Error
	if err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		c.JSON(controllers.Rsp(controllers.OK_INSERT_FAILED, nil, "用户名重复"))
		// c.JSON(http.StatusInternalServerError, controllers.SetRsp(controllers.ServerInternalDuplicateEntry, nil))
		return
	}
	if err != nil {
		utils.LogInstance().Println(err.Error())
		c.JSON(http.StatusInternalServerError, controllers.SetRsp(controllers.OK_INSERT_FAILED, nil))
		return
	}

	c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_INSERT_SUCCESS, admin))
	return
}