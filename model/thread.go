package model

import (
	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	ThreadId int    `gorm:"unique" json:"thread_id"` // check if needed to have string and not just int
	IsE      bool   `json:"is_e"`
	IsC      bool   `json:"is_c"`
	Brand    string `json:"brand"`
	Count    int    `json:"count"`
}
