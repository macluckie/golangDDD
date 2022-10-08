package api

import (
	h "test/golang/api/handler"
	docs "test/golang/docs"
	d "test/golang/domain"
	r "test/golang/repository"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(rp r.Repo, s d.Services) *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/v1"
	hand := h.NewHandler(rp, s)
	v1 := r.Group("/v1")
	{
		eg := v1.Group("/articles")
		{
			eg.GET("", hand.GetArticles)
			eg.GET("/:id", hand.GetArticle)
			eg.POST("", hand.CreateArticles)

		}
	}
	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
