package dto

import d "test/golang/domain"

type Article struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Author  string `json:"author" binding:"required"`
}

func ConverArticletoDomainArticle(article Article) d.Article {
	articleD := d.Article{}
	articleD.Autor = article.Author
	articleD.Content = article.Content
	articleD.Title = article.Title
	return articleD
}
