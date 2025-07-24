package handler

import (
	"gblog/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	BaseHandler
}

func NewArticleHandler(app *app.Application) *ArticleHandler {
	return &ArticleHandler{BaseHandler: BaseHandler{App: app}}
}

func (h *ArticleHandler) GetArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "article")
}

func (h *ArticleHandler) CreateArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Article created successfully"})
}
