package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/BON4/gosubs/config"
	models "github.com/BON4/gosubs/internal/domain/boil_postgres"

	herrors "github.com/BON4/gosubs/internal/errors"
	tokengen "github.com/BON4/gosubs/pkg/tokenGen"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	//authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	//authorizationPayloadKey = "authorization_payload"
)

type ServerMiddleware struct {
	tokenMaker tokengen.Generator
	headerKey  string
	payloadKey string
	logger     *logrus.Entry
}

func NewServerMiddleware(tgen tokengen.Generator, cfg config.ServerConfig, logger *logrus.Entry) *ServerMiddleware {
	return &ServerMiddleware{
		tokenMaker: tgen,
		headerKey:  cfg.HeaderKey,
		payloadKey: cfg.PaylaodKey,
		logger:     logger,
	}
}

func (m *ServerMiddleware) CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		ctx.Next()
	}
}

// AuthMiddleware creates a gin middleware for authorization
func (m *ServerMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(m.headerKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, herrors.ErrorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 1 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, herrors.ErrorResponse(err))
			return
		}

		accessToken := fields[0]
		payload, err := m.tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, herrors.ErrorResponse(err))
			return
		}

		ctx.Set(m.payloadKey, payload)
		ctx.Next()
	}
}

func (m *ServerMiddleware) RoleRestriction(alowedRoles ...models.AccountRole) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload, ok := ctx.Get(m.payloadKey)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, herrors.ErrorResponse(errors.New("TODO: custom error 1")))
			return
		}

		account, ok := payload.(*tokengen.Payload)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, herrors.ErrorResponse(errors.New("TODO: custom error 2")))
			return
		}

		var valid = false
		for _, role := range alowedRoles {
			if role == account.Instance.Role {
				valid = true
			}
		}

		if !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, herrors.ErrorResponse(errors.New("TODO: custom error 3")))
			return
		}

		ctx.Next()
	}
}
