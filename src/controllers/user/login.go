package user

import (
	"controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"models"
	"net/http"
	"time"
	"utils"
)

// curl "http://127.0.0.1:9088/admin/login" -d "name=lvyanpeng&password=123456"
func UserLogin(c *gin.Context) {
	keys := []string{"name", "password"}

	for _, v := range keys {
		if c.PostForm(v) == "" {
			c.JSON(http.StatusBadRequest, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "缺少参数:" + v,nil))
			return
		}
	}
	utils.LogInstance().Println(c.Request.PostForm)

	var user models.User
	user = models.User{
		Name: c.PostForm("name"),
	}
	db := models.Instance()
	err := db.Where(models.User{Name: c.PostForm("name")}).Not(models.User{Status: models.USER_ST_DELETED}).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "用户名错误", nil))
		return
	} else if err != nil {
		utils.LogInstance().Println(err.Error())
		c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_SERVER_ERROR, nil))
		return
	}
	if !utils.ChkPwd(user.Password, c.PostForm("password")) {
		c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "密码错误", nil))
		return
	}
	user.Token = utils.RandMd5()
	user.ExpiredAt = time.Now().Add(models.USER_TOKEN_DURATION)
	user.LastLoginAt = models.StdTime(time.Now())
	if err := db.Save(&user).Error; err != nil {
		utils.LogInstance().Println(err.Error())
		c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_SERVER_ERROR, nil))
		return
	}
	c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_INSERT_SUCCESS, user))
	return
}
