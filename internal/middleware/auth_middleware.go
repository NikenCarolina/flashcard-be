package middleware

import (
	"github.com/NikenCarolina/flashcard-be/internal/appconst"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO create authentication with JWT token
		ctx.Set(appconst.KeyUserID, 1)
		ctx.Next()
	}
}
