package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MyOrm *gorm.DB
var err error

func init() {
	MyOrm, err = gorm.Open("mysql", "root:123456789@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	//defer Db.Close()
	MyOrm.LogMode(true)
	if err != nil {
		panic("数据库连接失败!")
	}
	MyOrm.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	MyOrm.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	MyOrm.AutoMigrate(&User{}, &Article{})
}
