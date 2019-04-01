package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

//基础model, 别的表可以继承此表
type BaseModel struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt int64 `gorm:"not null"` //时间戳, unix, 秒 10位
	UpdatedAt int64 `gorm:"not null"`
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				nowTime := time.Now().Unix()
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				nowTime := time.Now().Unix()
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `UpdatedAt` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		nowTime := time.Now().Unix()
		scope.SetColumn("UpdatedAt", nowTime)
	}
}
