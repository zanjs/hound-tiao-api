package models

import (
	"time"
)

type (
	// IDModel is
	IDModel struct {
		ID uint `json:"id" sql:"AUTO_INCREMENT" gorm:"unique_index;not null;unique;primary_key;column:id"`
	}
	// UUIDModel is
	UUIDModel struct {
		UID string `json:"uid" sql:"index"  gorm:"unique_index;not null;unique;primary_key;column:uid"`
	}
)

type (
	// CreateModel is
	CreateModel struct {
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	}
	// UpdatedAtModel is
	UpdatedAtModel struct {
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
	}
	// DeletedAtModel is
	DeletedAtModel struct {
		DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	}
)

// BaseModel is
type BaseModel struct {
	IDModel
	UUIDModel
	CreateModel
	UpdatedAtModel
	DeletedAtModel
}

// IDBaseModel is
type IDBaseModel struct {
	IDModel
	CreateModel
	UpdatedAtModel
	DeletedAtModel
}

// UUIDBaseModel is
type UUIDBaseModel struct {
	UUIDModel
	CreateModel
	UpdatedAtModel
	DeletedAtModel
}

// IDCreateModel is
type IDCreateModel struct {
	IDModel
	CreateModel
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
