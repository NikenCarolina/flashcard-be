package middleware

import (
	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/dto"
	"github.com/gin-gonic/gin"
)

func Error() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if lastErr := ctx.Errors.Last(); lastErr != nil {
			switch e := lastErr.Err.(type) {
			case *apperror.Error:
				ctx.AbortWithStatusJSON(e.Code, dto.Response{Message: e.Error()})
			default:
				ctx.AbortWithStatusJSON(apperror.ErrInternalServerError.Code, dto.Response{Message: e.Error()})
			}
		}
	}
}
