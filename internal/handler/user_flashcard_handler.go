package handler

import (
	"net/http"

	"github.com/NikenCarolina/flashcard-be/internal/appconst"
	"github.com/NikenCarolina/flashcard-be/internal/apperror"
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

func (h *UserHandler) GetSetById(ctx *gin.Context) {
	var uri dto.FlashcardSetUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.Error(apperror.ErrBadRequest)
		return
	}

	data, err := h.useCase.GetSetById(ctx, ctx.GetInt(appconst.KeyUserID), uri.FlashcardSetID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: data,
	})
}

func (h *UserHandler) ListCards(ctx *gin.Context) {
	var uri dto.FlashcardSetUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.Error(apperror.ErrBadRequest)
		return
	}

	data, err := h.useCase.GetCards(ctx, ctx.GetInt(appconst.KeyUserID), uri.FlashcardSetID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: data,
	})
}
