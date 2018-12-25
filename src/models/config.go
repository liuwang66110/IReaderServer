package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gcfg.v1"
)

type cldConf struct {
	Server struct {
		Port string
	}
	Database struct {
		Dialect string
		Dsn     string
	}
}

var CldConf cldConf

func InitConf() {
	confFile := "./conf/" + gin.Mode() + ".ini"
	err := gcfg.ReadFileInto(&CldConf, confFile)
	if err != nil {
		panic("配置文件解析错误")
	}
	fmt.Println(CldConf)

}
