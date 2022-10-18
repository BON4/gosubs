package http

import (
	"net/http"
	"time"

	"github.com/BON4/gosubs/internal/domain"
	"github.com/BON4/gosubs/internal/errors"
	"github.com/BON4/gosubs/internal/server"
	"github.com/BON4/gosubs/internal/utis/tests"
	"github.com/gin-gonic/gin"
)

type accountHandler struct {
	userUc domain.AccountUsecase
	srv    *server.Server
}

func NewAccountHandler(g *gin.RouterGroup, uc domain.AccountUsecase, srv *server.Server) {
	handler := &accountHandler{
		userUc: uc,
		srv:    srv,
	}

	g.POST("/register", handler.Register)
}

type RegisterAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (t *accountHandler) Register(ctx *gin.Context) {
	req := RegisterAccountRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	domainAccount := domain.Account{
		Role:  domain.AccountRoleCreator,
		Email: req.Email,

		//TODO hash password
		Password: []byte(req.Password),
	}

	if err := t.userUc.Create(ctx.Request.Context(), &domainAccount); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

type loginAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type accountResponse struct {
	AccountID int64              `json:"account_id"`
	Email     string             `json:"email"`
	Role      domain.AccountRole `json:"role"`
}

type loginAccountResponse struct {
	AccessToken           string          `json:"access_token"`
	AccessTokenExpiresAt  time.Time       `json:"access_token_expires_at"`
	RefreshToken          string          `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time       `json:"refresh_token_expires_at"`
	Account               accountResponse `json:"account"`
}

func (t *accountHandler) Login(ctx *gin.Context) {
	var req loginAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errors.ErrorResponse(err))
		return
	}

	account, err := t.userUc.GetByEmail(ctx.Request.Context(), req.Email)
	if err != nil {
		//TODO if sql.ErrNoRows throw custom error
		ctx.JSON(http.StatusBadRequest, errors.ErrorResponse(err))
		return
	}

	err = tests.CheckPassword(req.Password, account.Password)
	if err != nil {
		//TODO throw custom error: Passwords dont match
		ctx.JSON(http.StatusUnauthorized, errors.ErrorResponse(err))
		return
	}

	acessToken, acessPayload, err := t.srv.Token.CreateToken(account, t.srv.Cfg.Token.AcessDuration)

	if err != nil {
		//TODO throw custom error: Passwords dont match
		ctx.JSON(http.StatusInternalServerError, errors.ErrorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := t.srv.Token.CreateToken(account, t.srv.Cfg.Token.RefreshDuration)
	if err != nil {
		//TODO throw custom error: Passwords dont match
		ctx.JSON(http.StatusInternalServerError, errors.ErrorResponse(err))
		return
	}

	if err := t.srv.Store.Set(ctx.Request.Context(), refreshPayload.ID.String(), &domain.Session{
		ID:           refreshPayload.ID,
		Instance:     *account,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIP:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}, t.srv.Cfg.Token.RefreshDuration); err != nil {
		//TODO throw custom error: Passwords dont match
		ctx.JSON(http.StatusInternalServerError, errors.ErrorResponse(err))
		return
	}

	resp := loginAccountResponse{
		AccessToken:           acessToken,
		AccessTokenExpiresAt:  acessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		Account: accountResponse{
			AccountID: account.AccountID,
			Email:     account.Email,
			Role:      account.Role,
		},
	}
	ctx.JSON(http.StatusOK, resp)
}
