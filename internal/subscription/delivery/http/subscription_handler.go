package http

import (
	"net/http"
	"time"

	"github.com/BON4/gosubs/config"
	"github.com/BON4/gosubs/internal/domain"
	herrors "github.com/BON4/gosubs/internal/errors"
	"github.com/BON4/gosubs/internal/middleware"
	tokengen "github.com/BON4/gosubs/pkg/tokenGen"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	null "github.com/volatiletech/null/v8"
)

type subscriptionHandler struct {
	subsUc domain.SubscriptionUsecase
	userUc domain.TgUserUsecase
	mid    *middleware.ServerMiddleware
	cfg    config.ServerConfig
	logger *logrus.Entry
}

func NewSubscriptionHandler(g *gin.RouterGroup, suc domain.SubscriptionUsecase, uuc domain.TgUserUsecase, mid *middleware.ServerMiddleware, cfg config.ServerConfig, logger *logrus.Entry) {

	handler := &subscriptionHandler{
		userUc: uuc,
		subsUc: suc,
		cfg:    cfg,
		mid:    mid,
		logger: logger,
	}

	g.POST("", mid.RoleRestriction(domain.AccountRoleAdmin, domain.AccountRoleBot), handler.Create)
	g.PATCH("", mid.RoleRestriction(domain.AccountRoleAdmin, domain.AccountRoleBot), handler.Update)
	g.GET("/list", handler.List)
}

// @Summary      List Subscriptions
// @Description  get subscription list. Only administrator and bot can get list of any accounts. Ordenery user can get list of subscriptions whitch belongs to his account.
// @Security     JWT
// @Tags         subscription
// @Accept       json
// @Produce      json
// @Param        input body   domain.FindSubRequest  true  "subscription list request filter"
// @Success      200     {array}   domain.Sub
// @Failure      400     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /sub/list [get]
func (t *subscriptionHandler) List(ctx *gin.Context) {
	req := domain.FindSubRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, herrors.ErrorResponse(err))
		return
	}

	payload, err := tokengen.GetPayloadFromContext(ctx, t.cfg.Auth.PaylaodKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, herrors.ErrorResponse(err))
		return
	}

	if payload.Instance.Role == domain.AccountRoleCreator {
		if req.AccountID != nil {
			if req.AccountID.Eq != payload.Instance.AccountID {
				ctx.Status(http.StatusMethodNotAllowed)
				return
			}
		}
	}

	subs, err := t.subsUc.List(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, herrors.ErrorResponse(err))
		return
	}

	//TODO: prepare sub list and dont returs domain struct
	ctx.JSON(http.StatusOK, subs)
}

// TODO: allow to update all subscription fields
type updateSubscriptionRequest struct {
	UserID    int64            `json:"user_id"`
	AccountID int64            `json:"account_id"`
	Status    domain.SubStatus `json:"status"`
}

// @Summary      Update
// @Description  updates subscription. Admin and bot can update subscription. Can be used to change subscription status, or price.
// @Security     JWT
// @Tags         subscription
// @Accept       json
// @Produce      json
// @Param        input body   updateSubscriptionRequest  true  "subscription new status and price"
// @Success      200
// @Failure      400     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /sub [patch]
func (t *subscriptionHandler) Update(ctx *gin.Context) {
	req := updateSubscriptionRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, herrors.ErrorResponse(err))
		return
	}

	sub, err := t.subsUc.GetByID(ctx.Request.Context(), req.UserID, req.AccountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, herrors.ErrorResponse(err))
		return
	}

	sub.Status = req.Status

	if err := t.subsUc.Update(ctx.Request.Context(), sub); err != nil {
		ctx.JSON(http.StatusInternalServerError, herrors.ErrorResponse(err))
		return
	}

	if _, err := t.subsUc.Save(ctx.Request.Context(), sub); err != nil {
		ctx.JSON(http.StatusInternalServerError, herrors.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusAccepted)
}

type createSubscriptionRequest struct {
	UserID    int64            `json:"user_id"`
	AccountID int64            `json:"account_id"`
	ExpiresAt time.Time        `json:"expires_at"`
	Price     int64            `json:"price"`
	Status    domain.SubStatus `json:"status"`
}

// @Summary      Create
// @Description  creates subscribtion with given users telegram_id and account_id. Only administrator and bot can create subscription
// @Security     JWT
// @Tags         subscription
// @Accept       json
// @Produce      json
// @Param        input body   createSubscriptionRequest  true  "subscription info"
// @Success      200
// @Failure      400     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /sub [post]
func (t *subscriptionHandler) Create(ctx *gin.Context) {
	req := createSubscriptionRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, herrors.ErrorResponse(err))
		return
	}

	user, err := t.userUc.GetByID(ctx.Request.Context(), req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, herrors.ErrorResponse(err))
		return
	}

	sub := &domain.Sub{
		UserID:      user.UserID,
		AccountID:   req.AccountID,
		ActivatedAt: time.Now(),
		ExpiresAt:   req.ExpiresAt,
		Status:      req.Status,
		Price:       null.IntFrom(int(req.Price)),
	}

	if err := t.subsUc.Create(ctx.Request.Context(), sub); err != nil {
		ctx.JSON(http.StatusInternalServerError, herrors.ErrorResponse(err))
		return
	}

	if _, err := t.subsUc.Save(ctx.Request.Context(), sub); err != nil {
		ctx.JSON(http.StatusInternalServerError, herrors.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
