package model

import (
	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	UserID      uint   `gorm:"uniqueIndex:idx_user_thread" json:"user_id"`
	User        User   `gorm:"foreignKey:UserID" json:"-"`
	ThreadId    string `gorm:"uniqueIndex:idx_user_thread" json:"thread_id"`
	IsE         bool   `json:"is_e"`
	IsC         bool   `json:"is_c"`
	Brand       string `json:"brand"`
	ThreadCount int64  `json:"thread_count"`
}

type ThreadDto struct {
	ThreadId    string `json:"thread_id"`
	IsE         bool   `json:"is_e"`
	IsC         bool   `json:"is_c"`
	Brand       string `json:"brand"`
	ThreadCount int64  `json:"thread_count"`
}

func (t *Thread) UpdateFields(other *Thread) {
	t.IsE = other.IsE
	t.IsC = other.IsC
	t.Brand = other.Brand
	t.ThreadCount = other.ThreadCount
}