package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	herrors "github.com/BON4/gosubs/internal/errors"
	tokengen "github.com/BON4/gosubs/pkg/tokenGen"
	"github.com/gin-gonic/gin"
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
}

func NewServerMiddleware(tgen tokengen.Generator, headerKey string, payloadKey string) *ServerMiddleware {
	return &ServerMiddleware{
		tokenMaker: tgen,
		headerKey:  headerKey,
		payloadKey: payloadKey,
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
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, herrors.ErrorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, herrors.ErrorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := m.tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, herrors.ErrorResponse(err))
			return
		}

		ctx.Set(m.payloadKey, payload)
		ctx.Next()
	}
}
