package models

import (
	"time"

	"anla.io/hound/db"
	"github.com/houndgo/suuid"
	gm "github.com/jinzhu/gorm"
)

type (
	// UserName is
	UserName struct {
		Username string `json:"username" gorm:"type:varchar(100);unique"`
	}
	// User is
	User struct {
		BaseModel
		UserName
		Email    string `json:"email" gorm:"type:varchar(100);unique"`
		Password string `json:"-"`
	}

	// UserShort is
	UserShort struct {
		UUIDModel
		UserName
	}
	// UserLogin is
	UserLogin struct {
		UserName
		Password string `json:"password"`
	}
)

//TableName is set User's table name to be `users`
func (UserShort) TableName() string {
	return "users"
}

//BeforeSave is
func (s *User) BeforeSave(scope *gm.Scope) (err error) {
	s.UID = suuid.New().String()
	return err
}

// Create is user
func (s User) Create(m *User) error {
	var (
		err error
	)
	m.CreatedAt = time.Now()
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// GetByUsername is find user
func (s User) GetByUsername(username string) (User, error) {
	var (
		user User
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&user, "username = ?", username).Error; err != nil {
		tx.Rollback()
		return user, err
	}
	tx.Commit()

	return user, err
}
