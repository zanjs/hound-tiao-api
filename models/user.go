package models

import (
	"time"

	"anla.io/hound/db"
)

type (
	// User is
	User struct {
		BaseModel
		Username string `json:"username" gorm:"type:varchar(100);unique"`
		Email    string `json:"email" gorm:"type:varchar(100);unique"`
		Password string `json:"-"`
	}
	// UserLogin is
	UserLogin struct {
		Username string `json:"username" gorm:"type:varchar(100);unique"`
		Password string `json:"password"`
	}
)

// Create is user
func (um User) Create(m *User) error {
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
func (um User) GetByUsername(username string) (User, error) {
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
