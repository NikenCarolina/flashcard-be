package handler

import (
	"net/http"

	"github.com/NikenCarolina/flashcard-be/internal/appconst"
	"github.com/NikenCarolina/flashcard-be/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) Register(ctx *gin.Context) {
	var req dto.User
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	err := h.useCase.Register(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Message: appconst.MsgRegisterOk,
	})

}

func (h *UserHandler) Login(ctx *gin.Context) {
	var req dto.User
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	token, err := h.useCase.Login(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("token", *token, 3600, "/", h.config.App.DomainName, true, true)
	ctx.JSON(http.StatusOK, dto.Response{
		Message: appconst.MsgLoginOk,
		Data:    dto.Redirect{URL: h.config.App.AuthRedirectURL},
	})
}

func (u *UserHandler) Logout(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("token", "", -1, "/", u.config.App.DomainName, true, true)

	ctx.JSON(http.StatusOK, dto.Response{
		Message: appconst.MsgLogoutOk,
	})
}

func (h *UserHandler) Profile(ctx *gin.Context) {
	data, err := h.useCase.Profile(ctx, ctx.GetInt(appconst.KeyUserID))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: appconst.MsgGetProfileOk,
		Data:    data,
	})
}
