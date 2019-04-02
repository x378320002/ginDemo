package models

type Article struct {
	BaseModel `mapstructure:",squash"` //squash是map转结构体时能解析继承的类用的(不能省略前面的逗号)
	Desc      string                   `gorm:"not null;defult:''"`
	ImgUrl    string                   `gorm:"not null;defult:''"`
	VideoUrl  string                   `gorm:"not null;defult:''"`
	ImgW      int                      `gorm:"not null;defult:0"`  //来自哪个网站
	ImgH      int                      `gorm:"not null;defult:0"`  //来自哪个网站
	Source    string                   `gorm:"not null;defult:''"` //来自哪个网站
	Type      int                      `gorm:"not null;defult:0"`  //什么类型的内容, 0, 图片+文字 1, gif+文字 2, 视频+文字 3, 文章
	//点赞数
	//收藏数
}

func ArticleList(list *[]Article) error {
	err := MyOrm.Find(list).Error
	return err
}
