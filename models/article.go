package models

type Article struct {
	BaseModel
	Desc   string `gorm:"not null;defult:''"`
	ImgUrl string `gorm:"not null;defult:0"`
	Source string `gorm:"not null;defult:0"` //来自哪个网站
	Type   int    `gorm:"not null;defult:0"` //什么类型的内容, 0, 图片+文字 1, gif+文字 2, 视频+文字 3, 文章
	//点赞数
	//收藏数
}

func ArticleList(list *[]Article) error {
	err := DbOrm.Find(list).Error
	return err
}
