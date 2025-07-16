package handler

import (
	"fmt"
	"gblog/app"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
	App *app.Application
}

func NewBaseHandler(app *app.Application) *BaseHandler {
	return &BaseHandler{App: app}
}

func (h *BaseHandler) RespondJSON(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, data)
}

func MakeRoutes(app *app.Application) {

	baseHandler := NewBaseHandler(app)
	articleHandler := NewArticleHandler(app)
	adminHandler := NewAdminHandler(app)
	apiHandler := NewAPIHandler(app)

	app.Router.GET("/", baseHandler.Home)

	articleGroup := app.Router.Group("/articles")
	{
		articleGroup.GET("/:id", articleHandler.GetArticle)
		articleGroup.POST("/", articleHandler.CreateArticle)
	}

	apiGroup := app.Router.Group("/api")
	{
		apiGroup.GET("/articles", apiHandler.Home)
	}

	adminGroup := app.Router.Group("/admin")
	{
		adminGroup.POST("/login", adminHandler.Login)
		adminGroup.GET("/articles", adminHandler.Articles)
	}
}

func (h *BaseHandler) Home(ctx *gin.Context) {
	count := h.App.DB.HgetInt("count", []byte("count"))
	h.App.DB.Hincr("count", []byte("count"), 1)
	h.RespondJSON(ctx, 200, fmt.Sprintf("Hello, World! %v", count))
}
