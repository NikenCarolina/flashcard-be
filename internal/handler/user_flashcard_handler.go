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
		Message: appconst.MsgListSetOk,
		Data:    data,
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
		Message: appconst.MsgGetSetOk,
		Data:    data,
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
		Message: appconst.MsgListCardOk,
		Data:    data,
	})
}

func (h *UserHandler) CreateCard(ctx *gin.Context) {
	var uri dto.FlashcardSetUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.Error(err)
		return
	}

	data, err := h.useCase.CreateCard(ctx, ctx.GetInt(appconst.KeyUserID), uri.FlashcardSetID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Message: appconst.MsgCreateCardOk,
		Data:    data,
	})
}

func (h *UserHandler) UpdateCard(ctx *gin.Context) {
	var uri dto.FlashcardUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.Error(err)
		return
	}

	var req dto.Flashcard
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}
	req.FlashcardID = int64(uri.FlashcardID)
	req.FlashcardSetID = int64(uri.FlashcardSetID)

	err := h.useCase.UpdateCard(ctx, ctx.GetInt(appconst.KeyUserID), &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Message: appconst.MsgUpdateCardOk,
	})
}

func (h *UserHandler) DeleteCard(ctx *gin.Context) {
	var uri dto.FlashcardUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.Error(err)
		return
	}

	err := h.useCase.DeleteCard(ctx, ctx.GetInt(appconst.KeyUserID), int(uri.FlashcardSetID), int(uri.FlashcardID))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: appconst.MsgDeleteCardOk,
	})
}

func (h *UserHandler) BulkUpdateCard(ctx *gin.Context) {
	var uri dto.FlashcardSetUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.Error(err)
		return
	}

	var flashcards []dto.Flashcard
	if err := ctx.ShouldBindJSON(&flashcards); err != nil {
		ctx.Error(err)
		return
	}

	var flashcardResponses []dto.FlashcardUpdateResponse
	for _, card := range flashcards {
		card.FlashcardSetID = int64(uri.FlashcardSetID)
		err := h.useCase.UpdateCard(ctx, ctx.GetInt(appconst.KeyUserID), &card)
		if err != nil {
			if serr, ok := err.(*apperror.Error); ok {
				flashcardResponses = append(flashcardResponses, dto.FlashcardUpdateResponse{
					Status:      int64(serr.Code),
					FlashcardID: card.FlashcardID,
				})
			} else {
				flashcardResponses = append(flashcardResponses, dto.FlashcardUpdateResponse{
					Status:      http.StatusInternalServerError,
					FlashcardID: card.FlashcardID,
				})
			}
		} else {
			flashcardResponses = append(flashcardResponses, dto.FlashcardUpdateResponse{
				Status:      http.StatusOK,
				FlashcardID: card.FlashcardID,
			})
		}

	}

	ctx.JSON(http.StatusMultiStatus, dto.Response{
		Message: appconst.MsgUpdateCardOk,
		Data:    flashcardResponses,
	})

}
