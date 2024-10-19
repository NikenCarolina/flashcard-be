package handler

import (
	"net/http"

	"github.com/NikenCarolina/flashcard-be/internal/appconst"
	"github.com/NikenCarolina/flashcard-be/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) ListSets(ctx *gin.Context) {
	data, err := h.useCase.GetSets(ctx, ctx.GetInt(appconst.KeyUserID))

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: data,
	})
}
