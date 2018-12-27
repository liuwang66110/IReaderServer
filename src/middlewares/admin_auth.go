package middlewares

import (
	"controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"models"
	"net/http"
	"time"
	"utils"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.LogInstance().Println(c.Request.URL)
		if c.PostForm("token") == "" {
			c.JSON(http.StatusBadRequest, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "缺少参数:token", nil))
			c.Abort()
			return
		}
		var user models.User
		user = models.User{
			Name: c.PostForm("token"),
		}
		db := models.Instance()
		err := db.Where(models.User{Token: c.PostForm("token")}).First(&user).Error
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusForbidden, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "token错误", nil))
			c.Abort()
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, controllers.SetRsp(controllers.OK_INSERT_FAILED, nil))
			c.Abort()
			return
		}
		if !user.ExpiredAt.After(time.Now()) {
			c.JSON(http.StatusForbidden, controllers.SetRspMsg(controllers.OK_TOKEN_FAILED, "token过期", nil))
			c.Abort()
			return
		}
		utils.LogInstance().Println(user)
		// TODO: token先不变更, 方便调试
		// admin.Token = utils.RandMd5()
		user.ExpiredAt = time.Now().Add(models.USER_TOKEN_DURATION)
		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, controllers.SetRsp(controllers.OK_INSERT_FAILED, nil))
		}
		// c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_INSERT_SUCCESS, admin))
		c.Set("user", user)
		c.Next()
	}
}
