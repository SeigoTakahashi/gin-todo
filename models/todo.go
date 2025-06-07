package models

import "gorm.io/gorm"

// Todoモデルの定義
type Todo struct {
	gorm.Model
	Task      string `gorm:"not null" json:"task" form:"task"`
	Completed bool   `gorm:"not null;default:false" json:"completed"`
}
