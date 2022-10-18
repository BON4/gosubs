package http

import (
	"net/http"

	"github.com/BON4/gosubs/internal/domain"
	"github.com/gin-gonic/gin"
)

type tgUserHandler struct {
	userUc domain.TgUserUsecase
}

func NewTgUserHandler(g *gin.RouterGroup, uc domain.TgUserUsecase) {
	handler := &tgUserHandler{
		userUc: uc,
	}

	g.POST("/tgusers", handler.Create)
}

func (t *tgUserHandler) Create(ctx *gin.Context) {

	ctx.JSON(http.StatusNotImplemented, gin.H{})
}
