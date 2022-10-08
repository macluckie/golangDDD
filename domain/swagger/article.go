package swagger

import "time"

type Article struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
	Title     string
	Autor     string
	Content   string
	Category  []Category `gorm:"many2many:article_category"`
}
