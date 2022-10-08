package domain

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string
	Autor    string
	Content  string
	// Category []Category `gorm:"many2many:article_category"`
}
