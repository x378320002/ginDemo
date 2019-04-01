package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DbOrm *gorm.DB
var err error

func init() {
	DbOrm, err = gorm.Open("mysql", "root:123456789@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	//defer Db.Close()
	DbOrm.LogMode(true)
	if err != nil {
		panic("数据库连接失败!")
	}
	DbOrm.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	DbOrm.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	DbOrm.AutoMigrate(&User{}, &Article{})
}
