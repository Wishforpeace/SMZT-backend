package jottings

import (
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type Jotting struct {
	gorm.Model
	StudentID string `gorm:column`
}
