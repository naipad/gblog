package handler

import (
	"gblog/app"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	BaseHandler
}

func NewAPIHandler(app *app.Application) *APIHandler {
	return &APIHandler{BaseHandler: BaseHandler{App: app}}
}

func (h *APIHandler) RespondJSON(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, data)
}
