package models

type (
	// Comment is articel comment
	Comment struct {
		BaseModel
		UserID    uint
		ArticleID uint   `json:"article_id"`
		Content   string `json:"content"`
		NickName  string `json:"nick_name" gorm:"-"`
		AvatarURL string `json:"avatar_url" gorm:"-"`
	}
)
