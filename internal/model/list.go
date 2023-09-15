package model

import "moonlay-todolist/internal/abstraction"

type List struct {
	abstraction.Entity
	Title       string     `json:"title" gorm:"size:100"`
	Description string     `json:"description" gorm:"size:1000"`
	Files       []ListFile `json:"files"`
	SubList     []SubList  `json:"sub_list"`
}

type ListFile struct {
	abstraction.Entity
	ListID string `json:"list_id" gorm:"foreignKey:ListID"`
	Link   string `json:"link"`
}
