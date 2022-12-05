package http

import (
	"net/http"
	"time"

	"github.com/BON4/gosubs/config"
	"github.com/BON4/gosubs/internal/domain"
	models "github.com/BON4/gosubs/internal/domain/boil_postgres"
	myerrors "github.com/BON4/gosubs/internal/errors"

	"github.com/BON4/gosubs/internal/middleware"
	"github.com/BON4/gosubs/internal/utis/tests"
	tokengen "github.com/BON4/gosubs/pkg/tokenGen"
	"github.com/BON4/timedQ/pkg/ttlstore"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type authHandler struct {
	accountUc domain.AccountUsecase
	mid       *middleware.ServerMiddleware
	cfg       config.ServerConfig
	logger    *logrus.Entry
	token     tokengen.Generator
	store     *ttlstore.MapStore[string, *domain.Session]
}

func NewAuthHandler(g *gin.RouterGroup, uc domain.AccountUsecase, mid *middleware.ServerMiddleware, cfg config.ServerConfig, token tokengen.Generator, store *ttlstore.MapStore[string, *domain.Session], logger *logrus.Entry) {
	handler := &authHandler{
		accountUc: uc,
		cfg:       cfg,
		mid:       mid,
		logger:    logger,
		token:     token,
		store:     store,
	}

	g.POST("/register", handler.Register)
	g.POST("/login", handler.Login)
}

// TODO: create registration/login via telegram username
type registerAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary      Register
// @Description  registers new account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input   body      registerAccountRequest  true  "register credantials"
// @Success      201
// @Failure      204     {object}  error
// @Failure      400     {object}  error
// @Failure      409     {object}  error
// @Failure      500     {object}  error
// @Router       /register [post]
func (t *authHandler) Register(ctx *gin.Context) {
	req := registerAccountRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	hashed, err := tests.HashPassword(req.Password)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	domainAccount := models.Account{
		Role:     models.AccountRoleCreator,
		Email:    req.Email,
		Password: hashed,
	}

	//TODO: Only for debug purpse. Figure out another way of creating administaror.
	if req.Email == "admin" {
		domainAccount.Role = models.AccountRoleAdmin
	}

	//TODO: Only for debug purpse. Figure out another way of creating bots.
	if req.Email == "bot" {
		domainAccount.Role = models.AccountRoleBot
	}

	if err := t.accountUc.Create(ctx.Request.Context(), &domainAccount); err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
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
	Role      models.AccountRole `json:"role"`
}

type loginAccountResponse struct {
	AccessToken           string          `json:"access_token"`
	AccessTokenExpiresAt  time.Time       `json:"access_token_expires_at"`
	RefreshToken          string          `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time       `json:"refresh_token_expires_at"`
	Account               accountResponse `json:"account"`
}

// @Summary      Login
// @Description  logins in to account with user provided credantials
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input   body      loginAccountRequest  true  "login credentials"
// @Success      200     {object}  loginAccountResponse
// @Failure      204     {object}  error
// @Failure      400     {object}  error
// @Failure      409     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /login [post]
func (t *authHandler) Login(ctx *gin.Context) {
	var req loginAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	account, err := t.accountUc.GetByEmail(ctx.Request.Context(), req.Email)
	if err != nil {
		//TODO if sql.ErrNoRows throw custom error
		ctx.JSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	err = tests.CheckPassword(req.Password, account.Password)
	if err != nil {
		//TODO throw custom error: Passwords dont match
		ctx.JSON(http.StatusUnauthorized, myerrors.ErrorResponse(err))
		return
	}

	acessToken, acessPayload, err := t.token.CreateToken(account, t.cfg.AcessDuration)

	if err != nil {
		//TODO throw custom error: Passwords dont match
		ctx.JSON(http.StatusInternalServerError, myerrors.ErrorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := t.token.CreateToken(account, t.cfg.RefreshDuration)
	if err != nil {
		//TODO throw custom error: Passwords dont match
		ctx.JSON(http.StatusInternalServerError, myerrors.ErrorResponse(err))
		return
	}

	if err := t.store.Set(ctx.Request.Context(), refreshPayload.ID.String(), &domain.Session{
		ID:           refreshPayload.ID,
		Instance:     *account,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIP:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}, t.cfg.RefreshDuration); err != nil {
		ctx.JSON(http.StatusInternalServerError, myerrors.ErrorResponse(err))
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
