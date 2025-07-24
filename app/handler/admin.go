package handler

import (
	"gblog/app"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	BaseHandler
}

func NewAdminHandler(app *app.Application) *AdminHandler {
	return &AdminHandler{BaseHandler: BaseHandler{App: app}}
}

func (h *AdminHandler) Login(ctx *gin.Context) {
	username := ctx.DefaultPostForm("username", "")
	password := ctx.DefaultPostForm("password", "")

	if username == "admin" && password == "password" {
		h.RespondJSON(ctx, http.StatusOK, "Login successful")
	} else {
		h.RespondJSON(ctx, http.StatusUnauthorized, "Invalid credentials")
	}
}

func (h *AdminHandler) Articles(ctx *gin.Context) {
	log.Println("Fetching articles list...")
	h.RespondJSON(ctx, http.StatusOK, "Articles list")
}
