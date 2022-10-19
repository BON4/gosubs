package http

import (
	"net/http"
	"strconv"

	"github.com/BON4/gosubs/internal/domain"
	herrors "github.com/BON4/gosubs/internal/errors"

	"github.com/BON4/gosubs/internal/server"
	tokengen "github.com/BON4/gosubs/pkg/tokenGen"
	"github.com/gin-gonic/gin"
)

type accountHandler struct {
	accountUc domain.AccountUsecase
	srv       *server.Server
}

func NewAccountHandler(g *gin.RouterGroup, uc domain.AccountUsecase, srv *server.Server) {
	handler := &accountHandler{
		accountUc: uc,
		srv:       srv,
	}

	g.GET("/list", srv.MidWar.RoleRestriction(domain.AccountRoleAdmin), handler.ListAccounts)

	g.GET("/:acc_id", handler.GetAccount)
	g.PATCH("/email", handler.UpdateEmail)
	g.DELETE("/:acc_id", srv.MidWar.RoleRestriction(domain.AccountRoleAdmin), handler.DeleteAccount)

}

func (t *accountHandler) GetAccount(ctx *gin.Context) {
	payload, err := tokengen.GetAccountFromContext(ctx, t.srv.Cfg.Auth.PaylaodKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, herrors.ErrorResponse(err))
		return
	}
	req_acc_id, err := strconv.ParseInt(ctx.Query("acc_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, herrors.ErrorResponse(err))
		return
	}

	if payload.Instance.Role != domain.AccountRoleAdmin {
		if req_acc_id != payload.Instance.AccountID {
			ctx.Status(http.StatusMethodNotAllowed)
			return
		}
	}

	acc, err := t.accountUc.GetByID(ctx.Request.Context(), req_acc_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, herrors.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, acc)
}

func (t *accountHandler) DeleteAccount(ctx *gin.Context) {
	acc_id, err := strconv.ParseInt(ctx.Query("acc_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, herrors.ErrorResponse(err))
		return
	}

	if err := t.accountUc.Delete(ctx.Request.Context(), int64(acc_id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, herrors.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

func (t *accountHandler) ListAccounts(ctx *gin.Context) {
	req := domain.FindAccountRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	accounts, err := t.accountUc.List(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, herrors.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type updateAccountEmailRequest struct {
	Email string `json:"email"`
}

func (t *accountHandler) UpdateEmail(ctx *gin.Context) {
	payload, err := tokengen.GetAccountFromContext(ctx, t.srv.Cfg.Auth.PaylaodKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, herrors.ErrorResponse(err))
		return
	}

	req := updateAccountEmailRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	acc, err := t.accountUc.GetByID(ctx.Request.Context(), payload.Instance.AccountID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, herrors.ErrorResponse(err))
		return
	}

	acc.Email = req.Email

	if err := t.accountUc.Update(ctx.Request.Context(), acc); err != nil {
		ctx.JSON(http.StatusInternalServerError, herrors.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusAccepted)
}
