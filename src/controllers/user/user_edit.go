package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"controllers"
	"models"
	"strconv"
	"fmt"
)

func UserEdit(c *gin.Context)  {
	_, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "请先登陆", nil))
		return
	}

	var name string
	var mobile string
	var email string
	userID, idError := strconv.ParseUint(c.PostForm("user_id"), 10, 64)
	if (idError != nil) {
		c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "id错误", nil))
		return
	}
	if (c.PostForm("name") != "") {
		tmpName := c.PostForm("name")
		name = tmpName
	}
	if (c.PostForm("mobile") != "") {
		tmpMobile := c.PostForm("mobile")
		mobile = tmpMobile
	}
	if (c.PostForm("email") != "") {
		tmpEmail := c.PostForm("email")
		email = tmpEmail
	}
	gender, _ := strconv.Atoi(c.PostForm("gender"))
	if gender > 0 && (gender != models.USER_GENDER_FEMALE && gender != models.USER_GENDER_MALE) {
		c.JSON(http.StatusOK, controllers.SetRspMsg(controllers.OK_INSERT_FAILED, "性别错误", nil))
		return
	}
	userInfo := models.UserInfo{
		Name: name,
		Mobile: mobile,
		Email: email,
		Gender: gender,
		UserID: userID,
	}
	db := models.Instance()
	if err := db.Model(models.UserInfo{}).Where(models.UserInfo{UserID: userID}).Updates(userInfo).Error; err != nil {
		c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_SERVER_ERROR, nil))
		return
	}
	fmt.Println(userInfo)
	c.JSON(http.StatusOK, controllers.SetRsp(controllers.OK_INSERT_SUCCESS, userInfo))
	return
}
