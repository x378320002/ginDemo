package models

//用户信息表
type User struct {
	BaseModel
	Name      string `gorm:"not null;size:50;unique_index"` //登陆的账号名
	Pwd       string `gorm:"not null;size:50"`
	Phone     string `gorm:"not null;index;size:50"`
	Email     string `gorm:"not null;index"`
	NickName  string `gorm:"not null;index"`
	HeaderUrl string `gorm:"not null"` //用户头像
	Age       int    `gorm:"not null"`
	Gender    byte   `gorm:"not null"`
	Desc      string `gorm:"not null;type:text"` //用户描述
}

func FindUserByName(user *User) bool {
	first := MyOrm.Where("name = ?", user.Name).First(user)
	if first.RecordNotFound() {
		return false
	}
	return true
}
