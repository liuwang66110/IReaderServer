package main

import (
	"fmt"
	"models"
	"github.com/gin-gonic/gin"
	"controllers/user"
	"controllers/friend"
	"middlewares"
)

func main() {
	fmt.Println("Cloud server start...")
	models.InitConf()

	initDB()
	router := initRouter()
	//绑定端口9088
	router.Run(":9088")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/v1/user/register", user.UserRegister)
	router.POST("/v1/user/login", user.UserLogin)
	router.POST("/v1/user/search", middlewares.AdminAuth(), user.UserSearch)
	router.POST("/v1/user/edit", middlewares.AdminAuth(), user.UserEdit)

	router.POST("/v1/friend/list", middlewares.AdminAuth(), friend.FriendList)
	return router
}

func initDB() {
	db := models.Instance()
	//defer db.Close()
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(300)

	// 初始化表
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.UserFriend{})
	db.AutoMigrate(&models.UserInfo{})
}
