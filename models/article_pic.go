package models

type (
	// ArticlePic 文字图片
	ArticlePic struct {
		ArticleID int    `json:"article_id"`
		Src       string `json:"src"`
	}
)
