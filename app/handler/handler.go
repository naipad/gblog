package handler

import (
	"embed"
	"gblog/app"
	"gblog/app/views"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
	App *app.Application
}

func NewBaseHandler(app *app.Application) *BaseHandler {
	return &BaseHandler{
		App: app,
	}
}

func (h *BaseHandler) RespondJSON(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, data)
}

func MakeRoutes(app *app.Application, assetFiles embed.FS, templateFiles embed.FS) {

	// bind templates
	_template, _ := template.ParseFS(views.TemplateFiles, "templates/**/*.html")
	app.Router.SetHTMLTemplate(_template)

	// bind assets
	assets, _ := fs.Sub(views.AssetFiles, "assets")
	app.Router.StaticFS("/assets", http.FS(assets))

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
	ctx.HTML(200, "home_index", gin.H{
		"Title": "Welcome to My Blog",
		"Count": count,
	})
}
