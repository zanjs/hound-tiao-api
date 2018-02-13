package models

import (
	"time"
)

// BaseModel is
type BaseModel struct {
	ID        int        `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key,column:id"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

// PageModel is
type PageModel struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Count  int `json:"count"`
}

// QueryParams is
type QueryParams struct {
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	WareroomID int    `json:"wareroom_id"`
	Day        int    `json:"day"`
	ProductID  int    `json:"product_id"`
}

// QueryParamsTime is
type QueryParamsTime struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// PathParams is
type PathParams struct {
	ID uint64 `json:"id"`
}
