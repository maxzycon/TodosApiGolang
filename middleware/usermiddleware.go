package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"todosAPI/src/token"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload_key"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizeHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizeHeader) == 0 {
			err := errors.New("authorization is not provide")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,err.Error())
			return
		}

		field := strings.Fields(authorizeHeader)
		if len(field) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,err.Error())
			return
		}

		authorizationType := strings.ToLower(field[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s",authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,err.Error())
			return
		}

		accessToken := field[1]
		payload,err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,err.Error())
			return
		}

		ctx.Set(authorizationPayloadKey,payload)
		ctx.Next()
	}
}