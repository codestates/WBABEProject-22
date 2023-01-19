package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"

	"oos/logger"
)

var (
	signingKey = []byte("secret")

	keyFunc = func(ctx context.Context) (interface{}, error) {
		return signingKey, nil
	}

	issuer = "go-jwt-middleware-example"

	audience = []string{"audience-example"}

	customClaims = func() validator.CustomClaims {
		return &CustomClaims{}
	}
)

type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func (c CustomClaims) HasScope(expectedScope string) bool {
	result := strings.Split(c.Scope, " ")

	for i := range result {
		if result[i] == expectedScope {
			return true
		}
	}

	return false
}

func ValidateToken() gin.HandlerFunc {
	jwtValidator, err := validator.New(
		keyFunc,
		validator.HS256,
		issuer,
		audience,
		validator.WithCustomClaims(customClaims),
		validator.WithAllowedClockSkew(30*time.Second),
	)
	if err != nil {
		logger.Fatal("failed to set up the validator: %v", err)
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		logger.Error("encountered error while validating JWT: %v", err)
	}

	jwtMiddleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return adapter.Wrap(jwtMiddleware.CheckJWT)
}

func ValidateScope(permission string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, ok := ctx.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		if !ok {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Failed to get validated JWT claims."},
			)
			return
		}

		customClaims, ok := claims.CustomClaims.(*CustomClaims)
		if !ok {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Failed to cast custom JWT claims to specific type."},
			)
			return
		}

		if !customClaims.HasScope(permission) {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				map[string]string{"message": "Insufficient scope."},
			)
			return
		}

		ctx.Next()
	}
}

// References
// https://github.com/auth0/go-jwt-middleware/tree/master/examples/gin-example
// https://dev.to/ksivamuthu/auth0-jwt-middleware-in-go-gin-web-framework-37mj
