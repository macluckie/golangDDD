package repository

import (
	"context"
	d "test/golang/domain"
)

func (r Repo) GetArticles(ctx context.Context) ([]d.Article, error) {
	articles := []d.Article{}
	r.R.Find(&articles)
	return articles, nil
}

func (r Repo) GetArticle(ctx context.Context, id int) (d.Article, error) {
	article := d.Article{}
	result := r.R.Where("id = ?", id).First(&article)
	if result.Error != nil {
		return d.Article{}, result.Error
	}
	return article, nil
}

func (r Repo) CreateArticle(ctx context.Context, article d.Article) (d.Article, error) {
	result := r.R.Create(&article)
	if result.Error != nil {
		return d.Article{}, result.Error
	}

	return article, nil
}
