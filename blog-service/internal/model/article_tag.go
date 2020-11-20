package model

type ArticleTag struct {
	*Model
	TagId     int `json:"tag_id"`
	ArticleId int `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
