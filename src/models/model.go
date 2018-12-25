package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sync"
	"utils"
)

var dao *gorm.DB
var one sync.Once

func Instance() *gorm.DB {
	one.Do(func() {
		var err error
		dao, err = gorm.Open(CldConf.Database.Dialect, CldConf.Database.Dsn)
		if err != nil {
			utils.LogInstance().Println("数据库连接失败: ", err)
			panic("数据库连接失败: " + err.Error())
		}
	})
	return dao
}
