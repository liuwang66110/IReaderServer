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
			c.JSON(http.StatusBadRequest, controllers.SetRsp(controllers.OK_INSERT_FAILED, nil))
			c.Abort()
			return
		}
		var admin models.User
		admin = models.User{
			Name: c.PostForm("token"),
		}
		db := models.Instance()
		err := db.Where(models.User{Token: c.PostForm("token")}).First(&admin).Error
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusForbidden, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "token错误", nil))
			c.Abort()
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, controllers.SetRsp(controllers.OK_INSERT_FAILED, nil))
			c.Abort()
			return
		}
		if !admin.ExpiredAt.After(time.Now()) {
			c.JSON(http.StatusForbidden, controllers.SetRspMsg(controllers.OK_TOKEN_FAILED, "token过期", nil))
			c.Abort()
			return
		}
		utils.LogInstance().Println(admin)
		// TODO: token先不变更, 方便调试
		// admin.Token = utils.RandMd5()
		admin.ExpiredAt = time.Now().Add(models.USER_TOKEN_DURATION)
		if err := db.Save(&admin).Error; err != nil {
			c.JSON(http.StatusInternalServerError, controllers.SetRsp(controllers.OK_INSERT_FAILED, nil))
		}
		// c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_INSERT_SUCCESS, admin))
		c.Set("admin", admin)
		c.Next()
	}
}
