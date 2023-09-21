package model

import "moonlay-todolist/internal/abstraction"

type SubList struct {
	abstraction.Entity
	ListID      string        `json:"list_id"`
	Title       string        `json:"title" gorm:"size:100"`
	Description string        `json:"description" gorm:"size:1000"`
	Files       []SubListFile `json:"files"`
}

type SubListFile struct {
	abstraction.Entity
	SubListID string `json:"sub_list_id" gorm:"foreignKey:SubListID"`
	Link      string `json:"link"`
}
