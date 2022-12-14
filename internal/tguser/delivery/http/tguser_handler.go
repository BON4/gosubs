package http

import (
	"net/http"
	"strconv"

	"github.com/BON4/gosubs/config"
	"github.com/BON4/gosubs/internal/domain"
	models "github.com/BON4/gosubs/internal/domain/boil_postgres"
	myerrors "github.com/BON4/gosubs/internal/errors"
	"github.com/BON4/gosubs/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type tgUserHandler struct {
	userUc domain.TgUserUsecase
	mid    *middleware.ServerMiddleware
	cfg    config.ServerConfig
	logger *logrus.Entry
}

func NewTgUserHandler(g *gin.RouterGroup, uc domain.TgUserUsecase, mid *middleware.ServerMiddleware, cfg config.ServerConfig, logger *logrus.Entry) {
	handler := &tgUserHandler{
		userUc: uc,
		cfg:    cfg,
		mid:    mid,
		logger: logger,
	}

	g.GET("/list", handler.ListUsers)

	g.GET("/:usr_id", handler.GetUser)
	g.POST("", mid.RoleRestriction(models.AccountRoleAdmin, models.AccountRoleBot), handler.Create)
	g.PATCH("/:usr_id", handler.Update)
	g.DELETE("/:usr_id", mid.RoleRestriction(models.AccountRoleAdmin))
}

type updateUserRequest struct {
	Username string            `json:"username"`
	Status   models.UserStatus `json:"status"`
}

type updateUserResponse struct {
	UserID     int64             `json:"user_id"`
	TelegramID int64             `json:"telegram_id"`
	Username   string            `json:"username"`
	Status     models.UserStatus `json:"status"`
}

// @Summary      Update User
// @Description  updateds users username and/or status, provided by id. Admin and bot can update any user.
// @Security     JWT
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        usr_id path int64 true "user id"
// @Param        input body   updateUserRequest  true  "user new status or username"
// @Success      200     {object}  updateUserResponse
// @Failure      204     {object}  error
// @Failure      400     {object}  error
// @Failure      409     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /user/{usr_id}/password [patch]
func (t *tgUserHandler) Update(ctx *gin.Context) {
	req := updateUserRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	req_usr_id, err := strconv.ParseInt(ctx.Param("usr_id"), 10, 64)
	if err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	usr, err := t.userUc.GetByID(ctx.Request.Context(), req_usr_id)
	if err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	usr.Username = req.Username
	if req.Status != "" {
		usr.Status = req.Status
	}

	if err := t.userUc.Update(ctx.Request.Context(), usr); err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updateUserResponse{
		UserID:     usr.UserID,
		TelegramID: usr.TelegramID,
		Username:   usr.Username,
		Status:     usr.Status,
	})
}

// @Summary      List Users
// @Description  get user list. Only administrator and bot can get list of accounts
// @Security     JWT
// @Tags         user
// @Produce      json
// @Param        page_size         query     int              true "page size"
// @Param        page_number       query     int              true "page number"
// @Param        role_eq           query     string           false "role name is equal to"
// @Param        role_like         query     string           false "role name is like"
// @Param        username_eq       query     string           false "username is equal to"
// @Param        username_like     query     string           false "username is like"
// @Success      200     {array}   models.Tguser
// @Failure      204     {object}  error
// @Failure      400     {object}  error
// @Failure      409     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /user/list [get]
func (t *tgUserHandler) ListUsers(ctx *gin.Context) {
	req, err := domain.ParseFindUserRequest(ctx.Request.URL.Query())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	users, err := t.userUc.List(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	//TODO: do not respond with models.Tguser
	ctx.JSON(http.StatusOK, users)
}

// @Summary      Delete User
// @Description  deletes user object by given id
// @Security     JWT
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        usr_id path   int64  true  "user id"
// @Success      200
// @Failure      204     {object}  error
// @Failure      400     {object}  error
// @Failure      409     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /user/{usr_id} [delete]
func (t *tgUserHandler) DeleteUser(ctx *gin.Context) {
	usr_id, err := strconv.ParseInt(ctx.Param("usr_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	if err := t.userUc.Delete(ctx.Request.Context(), int64(usr_id)); err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

type getUserResponse struct {
	UserID     int64             `json:"user_id"`
	TelegramID int64             `json:"telegram_id"`
	Username   string            `json:"username"`
	Status     models.UserStatus `json:"status"`
}

// @Summary      Get User
// @Description  returns user object by given id
// @Security     JWT
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        usr_id path   int64  true  "user id"
// @Success      200     {object}  getUserResponse
// @Failure      204     {object}  error
// @Failure      400     {object}  error
// @Failure      409     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /user/{usr_id} [get]
func (t *tgUserHandler) GetUser(ctx *gin.Context) {
	req_usr_id, err := strconv.ParseInt(ctx.Param("usr_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	usr, err := t.userUc.GetByID(ctx.Request.Context(), req_usr_id)
	if err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, getUserResponse{
		UserID:     usr.UserID,
		TelegramID: usr.TelegramID,
		Username:   usr.Username,
		Status:     usr.Status,
	})
}

type createUserRequest struct {
	TelegramID int64             `json:"telegram_id"`
	Username   string            `json:"username"`
	Status     models.UserStatus `json:"status"`
}

type createUserResponse struct {
	UserID     int64             `json:"user_id"`
	TelegramID int64             `json:"telegram_id"`
	Username   string            `json:"username"`
	Status     models.UserStatus `json:"status"`
}

// @Summary      Create
// @Description  creates user with given users telegram_id and username. Only administrator and bot can create user
// @Security     JWT
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        input body   createUserRequest  true  "user info"
// @Success      200     {object}  createUserResponse
// @Failure      204     {object}  error
// @Failure      400     {object}  error
// @Failure      409     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /user [post]
func (t *tgUserHandler) Create(ctx *gin.Context) {
	req := createUserRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, myerrors.ErrorResponse(err))
		return
	}

	usr := &models.Tguser{
		TelegramID: req.TelegramID,
		Username:   req.Username,
		Status:     req.Status,
	}

	if err := t.userUc.Create(ctx.Request.Context(), usr); err != nil {
		ctx.JSON(myerrors.GetErrorCode(err), myerrors.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, createUserResponse{
		UserID:     usr.UserID,
		TelegramID: usr.TelegramID,
		Username:   usr.Username,
		Status:     usr.Status,
	})
}
