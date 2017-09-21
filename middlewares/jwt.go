package middlewares

import (
	"github.com/fpay/gopress"
	"github.com/labstack/echo/middleware"
	//"strings"
)

// NewJwtMiddleware returns jwt middleware.
func NewJwtMiddleware() gopress.MiddlewareFunc {
	err := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		TokenLookup: "query:token",
		Skipper: func(c gopress.Context) bool {
			//if strings.HasPrefix(c.Path(), "/users/login") {
			//	return true
			//}
			return false
		},
	})

	return err
}
