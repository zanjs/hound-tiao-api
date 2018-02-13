package models

import (
	"time"

	"anla.io/hound/db"
)

type (
	// Article is
	Article struct {
		BaseModel
		User    User `gorm:"ForeignKey:UserId;AssociationForeignKey:Id" json:"user"`
		Pics    []ArticlePic
		UserID  int    `json:"user_id" gorm:"type:integer(11)"`
		Title   string `json:"title" gorm:"type:varchar(100)"`
		Content string `json:"content" gorm:"type:text"`
	}
)

// Create is
func (a Article) Create(m *Article) error {
	var err error

	m.CreatedAt = time.Now()
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
