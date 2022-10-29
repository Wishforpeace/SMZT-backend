package jottings

import (
	"SMZT/model"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type Jotting struct {
	gorm.Model
	StudentID string `gorm:"column:student_id"`
	Title     string `gorm:"column:title"`
	Content   string `gorm:"column:content;size:1000"`
}

func GetJottings(student_id string) ([]Jotting, error) {
	var jot []Jotting
	if err := model.DB.Self.Model(&Jotting{}).Where("student_id = ?", student_id).Find(&jot).Error; err != nil {
		return nil, err
	}

	return jot, nil
}

func CreateJotting(student_id string, title string, content string) error {
	var jot = Jotting{
		StudentID: student_id,
		Title:     title,
		Content:   content,
	}

	if err := model.DB.Self.Create(&jot).Error; err != nil {
		return err
	}
	return nil
}
