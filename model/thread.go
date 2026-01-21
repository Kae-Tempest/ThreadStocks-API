package model

import (
	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	User     User   `gorm:"foreignkey:ID"`
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

func (t *Thread) UpdateFields(other *Thread) {
	t.IsE = other.IsE
	t.IsC = other.IsC
	t.Brand = other.Brand
	t.Count = other.Count
}
