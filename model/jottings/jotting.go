package jottings

import (
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type Jotting struct {
	gorm.Model
	StudentID string `gorm:"column:student_id"`
	Title     string `gorm:"column:title"`
	Content   string `gorm:"column:content;size:1000"`
}
