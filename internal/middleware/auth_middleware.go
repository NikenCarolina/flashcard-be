package middleware

import (
	"log"
	"net/http"

	"github.com/NikenCarolina/flashcard-be/internal/appconst"
	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/dto"
	"github.com/NikenCarolina/flashcard-be/internal/util"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwt util.JwtProvider
}

func NewAuthMiddleware(jwt util.JwtProvider) *AuthMiddleware {
	return &AuthMiddleware{
		jwt: jwt,
	}
}

func (m *AuthMiddleware) IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.SetSameSite(http.SameSiteNoneMode)

		token, err := ctx.Cookie("token")
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Message: apperror.ErrUnauthorized.Message,
			})
			return
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Message: apperror.ErrUnauthorized.Message,
			})
			return
		}

		claims, err := m.jwt.Parse(token)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Message: apperror.ErrUnauthorized.Message,
			})
			return
		}

		ctx.Set(appconst.KeyUserID, int(claims.UserID))
		ctx.Next()
	}
}
