package handler

import (
	"context"
	"strconv"

	"gorm.io/gorm"

	dto "test/golang/api/handler/dto"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1

// get Articles
// @Summary get Article
// @Schemes
// @Description get all articles
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} swagger.Article
// @Router /articles [get]
func (h Handler) GetArticles(g *gin.Context) {
	ctx := context.Background()
	h.Service.ArticleService.R = h.Rp
	articles, err := h.Service.ArticleService.GetArticlesService(ctx)
	if err != nil {
		g.JSON(400, gin.H{"error": err})
	}
	g.JSON(200, gin.H{"articles": articles})
}

func (h Handler) CreateArticles(g *gin.Context) {
	ctx := context.Background()
	h.Service.ArticleService.R = h.Rp
	sa := h.Service.ArticleService
	article := dto.Article{}
	err := g.ShouldBindJSON(&article)
	if err != nil {
		g.JSON(400, gin.H{"error json body": err})
	}
	articleD := dto.ConverArticletoDomainArticle(article)
	articleResult, er := sa.CreateArticleService(ctx, articleD)
	if er != nil {
		g.JSON(400, gin.H{"error json body": er})
	}
	g.JSON(200, gin.H{"article": articleResult})

}

func (h Handler) GetArticle(g *gin.Context) {
	ctx := context.Background()
	h.Service.ArticleService.R = h.Rp
	id, er := strconv.Atoi(g.Param("id"))
	if er != nil {
		g.JSON(400, gin.H{"error parse id": er})
	}
	article, err := h.Service.ArticleService.GetArticleService(ctx, id)
	if err == gorm.ErrRecordNotFound {
		g.JSON(404, gin.H{"article not found with Id": g.Param("id")})
		return
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		g.JSON(400, gin.H{"error": err})
		return
	}

	g.JSON(200, gin.H{"articles": article})
}
