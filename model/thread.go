package model

import (
	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	UserID   User   `gorm:"foreignkey:UserID"`
	ThreadId string `gorm:"unique" json:"thread_id"`
	IsE      bool   `json:"is_e"`
	IsC      bool   `json:"is_c"`
	Brand    string `json:"brand"`
	Count    int    `json:"count"`
}

type ThreadDto struct {
	ThreadId string `json:"thread_id"`
	IsE      bool   `json:"is_e"`
	IsC      bool   `json:"is_c"`
	Brand    string `json:"brand"`
}
