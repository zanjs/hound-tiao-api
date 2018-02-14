package models

import (
	"time"

	"anla.io/hound/db"
	"github.com/houndgo/suuid"
	gm "github.com/jinzhu/gorm"
)

type (
	// Article is
	Article struct {
		UUIDBaseModel
		User      UserShort    `gorm:"Table:user;ForeignKey:UserId;AssociationForeignKey:Id" json:"user"`
		Pics      []ArticlePic `json:"pics"`
		UserID    uint         `json:"user_id" gorm:"type:integer(11)"`
		Title     string       `json:"title" gorm:"type:varchar(100)"`
		Content   string       `json:"content" gorm:"type:text"`
		ViewCount int          `json:"view_count"`
		Disabled  bool         `json:"disabled" gorm:"default:'0'"`
		Comments  []*Comment   `json:"comments" gorm:"-"`
	}
)

//BeforeSave is
func (a *Article) BeforeSave(scope *gm.Scope) (err error) {
	a.UID = suuid.New().String()
	return err
}

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

// GetAll is find
func (a Article) GetAll() ([]Article, error) {
	var (
		data []Article
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Preload("User").Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()

	return data, err
}
