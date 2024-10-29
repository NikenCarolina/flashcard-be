package handler

import (
	"net/http"

	"github.com/NikenCarolina/flashcard-be/internal/appconst"
	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) StartSession(ctx *gin.Context) {
	var req dto.FlashcardSetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(apperror.ErrBadRequest)
		return
	}

	data, err := h.useCase.StartSession(ctx, ctx.GetInt(appconst.KeyUserID), req.FlashcardSetID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: appconst.MsgStartSessionOk,
		Data:    data,
	})
}

func (h *UserHandler) EndSession(ctx *gin.Context) {
	var uri dto.SessionUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.Error(apperror.ErrBadRequest)
		return
	}

	var req dto.EndSessionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(apperror.ErrBadRequest)
		return
	}

	err := h.useCase.EndSession(ctx, ctx.GetInt(appconst.KeyUserID), uri.SessionID, req.SetID, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: appconst.MsgEndSessionOk,
	})
}
