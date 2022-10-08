package domain

import "context"

type Repo interface {
	CreateArticle(ctx context.Context, article Article) (Article, error)
	GetArticle(ctx context.Context, id int) (Article, error)
	GetArticles(ctx context.Context) ([]Article, error)
}

type ArticleService struct {
	R Repo
}

func (a ArticleService) CreateArticleService(ctx context.Context, article Article) (*Article, error) {
	article, err := a.R.CreateArticle(ctx, article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (a ArticleService) GetArticleService(ctx context.Context, id int) (*Article, error) {
	article, err := a.R.GetArticle(ctx, id)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (a ArticleService) GetArticlesService(ctx context.Context) ([]Article, error) {
	article, err := a.R.GetArticles(ctx)
	if err != nil {
		return nil, err
	}
	return article, nil
}
