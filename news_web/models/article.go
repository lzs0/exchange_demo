package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `bind:"required"`
	Content string `bind:"required"`
	Preview string `bind:"required"`
}
