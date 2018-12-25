package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"controllers"
	"fmt"
)

func UserRegister(c *gin.Context)  {
	fmt.Println("register")
	c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_SUCCESS, "注册成功", nil))
	return
}