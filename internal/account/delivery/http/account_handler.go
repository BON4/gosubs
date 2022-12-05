package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/BON4/gosubs/config"
	"github.com/BON4/gosubs/internal/domain"
	models "github.com/BON4/gosubs/internal/domain/boil_postgres"
	myerrors "github.com/BON4/gosubs/internal/errors"
	"github.com/BON4/gosubs/internal/middleware"
	"github.com/BON4/gosubs/internal/utis/tests"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"

	tokengen "github.com/BON4/gosubs/pkg/tokenGen"
	"github.com/gin-gonic/gin"
)

type accountHandler struct {
	userUc    domain.TgUserUsecase
	accountUc domain.AccountUsecase
	mid       *middleware.ServerMiddleware
	cfg       config.ServerConfig
	logger    *logrus.Entry
}

func NewAccountHandler(g *gin.RouterGroup, uc domain.AccountUsecase, userUc domain.TgUserUsecase, mid *middleware.ServerMiddleware, cfg config.ServerConfig, logger *logrus.Entry) {
	handler := &accountHandler{
		accountUc: uc,
		userUc:    userUc,
		mid:       mid,
		cfg:       cfg,
		logger:    logger,
	}

	//TODO: maby allow json request along side with query params
	g.GET("/list", mid.RoleRestriction(models.AccountRoleAdmin), handler.ListAccounts)

	g.GET("/:acc_id", handler.GetAccount)
	g.PATCH("/:acc_id/email", handler.UpdateEmail)
	g.PATCH("/:acc_id/user", handler.UpdateUser)
	g.PATCH("/:acc_id/password", handler.UpdatePassword)
	g.DELETE("/:acc_id", mid.RoleRestriction(models.AccountRoleAdmin), handler.DeleteAccount)
}

// @Summary      Get Account
// @Description  get account by id. Creator can get only his account. Administrator can get any account
// @Security     JWT
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        acc_id   path      int64  true  "account id"
// @Success      200     {object}  models.Account
// @Failure      204     {object}  error
// @Failure      400     {object}  error
// @Failure      409     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /account/{acc_id} [get]
func (t *accountHandler) GetAccount(ctx *gin.Context) {
	payload, err := tokengen.GetPayloadFromContext(ctx, t.cfg.PaylaodKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, myerrors.ErrorResponse(err))
		return
	}

	req_acc_id, err := strconv.ParseInt(ctx.Param("acc_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	if payload.Instance.Role != models.AccountRoleAdmin {
		if req_acc_id != payload.Instance.AccountID {
			ctx.Status(http.StatusMethodNotAllowed)
			return
		}
	}

	acc, err := t.accountUc.GetByID(ctx.Request.Context(), req_acc_id)
	if err != nil {

		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	//TODO: do not return models.Account. Prepare it: hide password etc.
	ctx.JSON(http.StatusOK, acc)
}

// @Summary      Delete Account
// @Description  deletes an account. Only administrator can delete accounts.
// @Security     JWT
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        acc_id   path      int64  true  "account id"
// @Success      200
// @Failure      204     {object}  error
// @Failure      400     {object}  error
// @Failure      409     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /account/{acc_id} [delete]
func (t *accountHandler) DeleteAccount(ctx *gin.Context) {
	acc_id, err := strconv.ParseInt(ctx.Param("acc_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	if err := t.accountUc.Delete(ctx.Request.Context(), int64(acc_id)); err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary      List Accounts
// @Description  get account list. Only administrator can get list of accounts
// @Security     JWT
// @Tags         account
// @Produce      json
// @Param        page_size         query     int              true "page size"
// @Param        page_number       query     int              true "page number"
// @Param        status_eq         query     string           false "status name is equal to"
// @Param        status_like       query     string           false "status name is like"
// @Success      200     {array}   models.Account
// @Failure      204     {object}  error
// @Failure      400     {object}  error
// @Failure      409     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /account/list [get]
func (t *accountHandler) ListAccounts(ctx *gin.Context) {
	req, err := domain.ParseFindAccountRequest(ctx.Request.URL.Query())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	accounts, err := t.accountUc.List(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	//TODO: do not respond with models.Account
	ctx.JSON(http.StatusOK, accounts)
}

type updateAccountPasswordRequest struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password,omitempty"`
}

// @Summary      Update Password
// @Description  updates password for current account. Admin can change password without provieding an old password. Admin can update password for any user.
// @Security     JWT
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        acc_id path int64 true "account id"
// @Param        input body   updateAccountPasswordRequest  true  "account old and new password"
// @Success      200
// @Failure      204     {object}  error
// @Failure      400     {object}  error
// @Failure      409     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /account/{acc_id}/password [patch]
func (t *accountHandler) UpdatePassword(ctx *gin.Context) {
	payload, err := tokengen.GetPayloadFromContext(ctx, t.cfg.PaylaodKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	req_acc_id, err := strconv.ParseInt(ctx.Param("acc_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	if payload.Instance.Role != models.AccountRoleAdmin {
		if req_acc_id != payload.Instance.AccountID {
			ctx.Status(http.StatusMethodNotAllowed)
			return
		}
	}

	req := updateAccountPasswordRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	acc, err := t.accountUc.GetByID(ctx.Request.Context(), req_acc_id)
	if err != nil {

		ctx.JSON(http.StatusNotFound, myerrors.ErrorResponse(err))
		return
	}
	if payload.Instance.Role != models.AccountRoleAdmin {
		if err := tests.CheckPassword(req.OldPassword, acc.Password); err != nil {
			ctx.JSON(http.StatusBadRequest, myerrors.ErrorResponse(errors.New("Passwords dont match.")))
			return
		}
	}

	hashed, err := tests.HashPassword(req.NewPassword)

	acc.Password = hashed

	if err := t.accountUc.Update(ctx.Request.Context(), acc); err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusAccepted)
}

type updateAccountEmailRequest struct {
	Email string `json:"email"`
}

// @Summary      Update Email
// @Description  updates email for current user. Admin can update email for any user.
// @Security     JWT
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        acc_id path int64 true "account id"
// @Param        input body   updateAccountEmailRequest  true  "account new email"
// @Success      200
// @Failure      204     {object}  error
// @Failure      400     {object}  error
// @Failure      409     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /account/{acc_id}/email [patch]
func (t *accountHandler) UpdateEmail(ctx *gin.Context) {
	payload, err := tokengen.GetPayloadFromContext(ctx, t.cfg.PaylaodKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	req_acc_id, err := strconv.ParseInt(ctx.Param("acc_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	if payload.Instance.Role != models.AccountRoleAdmin {
		if req_acc_id != payload.Instance.AccountID {
			ctx.Status(http.StatusMethodNotAllowed)
			return
		}
	}

	req := updateAccountEmailRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	acc, err := t.accountUc.GetByID(ctx.Request.Context(), req_acc_id)
	if err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	acc.Email = req.Email

	if err := t.accountUc.Update(ctx.Request.Context(), acc); err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusAccepted)
}

type updateAccountUserRequest struct {
	UserID     int `json:"user_id"`
	TelegramID int `json:"telegram_id"`
}

// @Summary      Update TgUser
// @Description  updates telegram user conected to this account. Admin, can update email for any user. Either of one of the fields must be provided
// @Security     JWT
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        acc_id path int64 true "account id"
// @Param        input body   updateAccountUserRequest  true  "account new email"
// @Success      200
// @Failure      204     {object}  error
// @Failure      400     {object}  error
// @Failure      409     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /account/{acc_id}/user [patch]
func (t *accountHandler) UpdateUser(ctx *gin.Context) {
	payload, err := tokengen.GetPayloadFromContext(ctx, t.cfg.PaylaodKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	req_acc_id, err := strconv.ParseInt(ctx.Param("acc_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	if payload.Instance.Role != models.AccountRoleAdmin {
		if req_acc_id != payload.Instance.AccountID {
			ctx.Status(http.StatusMethodNotAllowed)
			return
		}
	}

	req := updateAccountUserRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	usr := &models.Tguser{}

	if req.TelegramID == 0 {
		usr, err = t.userUc.GetByID(ctx, int64(req.UserID))
		if err != nil {
			ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
			return
		}
	} else {
		usr, err = t.userUc.GetByTelegramID(ctx, int64(req.UserID))
		if err != nil {
			ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
			return
		}
	}

	acc, err := t.accountUc.GetByID(ctx.Request.Context(), req_acc_id)
	if err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	acc.UserID = null.Int64From(usr.UserID)

	if err := t.accountUc.Update(ctx.Request.Context(), acc); err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusAccepted)
}
