package models

import (
	"time"

	"anla.io/hound/db"
	"github.com/houndgo/suuid"
	"github.com/theplant/batchputs"
)

type (
	// Article is
	Article struct {
		UUIDBaseModel
		User      User         `gorm:"Table:user;ForeignKey:UserId;AssociationForeignKey:Id" json:"user,omitempty"`
		Pics      []ArticlePic `json:"pics,omitempty"`
		UserID    uint         `json:"user_id" gorm:"type:integer(11)"`
		Title     string       `json:"title" gorm:"type:varchar(100)"`
		Content   string       `json:"content" gorm:"type:text"`
		ViewCount int          `json:"view_count"`
		Disabled  bool         `json:"disabled" gorm:"default:'0'"`
		Comments  []*Comment   `json:"comments" gorm:"-"`
	}
)

//BeforeSave is
// func (a *Article) BeforeSave(scope *gm.Scope) (err error) {
// 	a.UID = suuid.New().String()
// 	return err
// }

// Create is
func (a Article) Create(m *Article) error {
	var err error
	m.UID = suuid.New().String()
	rows := [][]interface{}{}

	pics := m.Pics
	createTime := time.Now()
	for i := 0; i < len(pics); i++ {
		item := pics[i]
		rows = append(rows, []interface{}{
			m.UID,
			createTime,
			item.Src,
		})
	}

	columns := []string{"article_id", "created_at", "src"}
	dialect := "mysql"

	err = batchputs.Put(gorm.MysqlConn().DB(), dialect, "article_pics", "article_id", columns, rows)
	if err != nil {
		panic(err)
	}

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
