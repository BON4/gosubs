package http

import (
	"net/http"

	"github.com/BON4/gosubs/internal/domain"
	"github.com/BON4/gosubs/internal/errors"
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

type createUserRequest struct {
	TelegramID int64             `json:"telegram_id"`
	Username   string            `json:"username"`
	Status     domain.UserStatus `json:"status"`
}

type createUserResponse struct {
	UserID     int64             `json:"user_id"`
	TelegramID int64             `json:"telegram_id"`
	Username   string            `json:"username"`
	Status     domain.UserStatus `json:"status"`
}

func (t *tgUserHandler) Create(ctx *gin.Context) {
	req := &createUserRequest{}
	if err := ctx.BindJSON(req); err != nil {
		//TODO if sql.ErrNoRows throw custom error
		ctx.JSON(http.StatusBadRequest, errors.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusNotImplemented, gin.H{})
}
