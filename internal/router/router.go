package router

import (
	"net/http"

	"github.com/NikenCarolina/flashcard-be/internal/handler"
	"github.com/NikenCarolina/flashcard-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Init(opts *handler.HandlerOpts) http.Handler {
	r := gin.New()
	r.ContextWithFallback = true

	middlewares := []gin.HandlerFunc{
		middleware.Error(),
	}
	r.Use(middlewares...)

	authMiddleware := middleware.NewAuthMiddleware()

	r.GET("/sets", authMiddleware.IsAuthenticated(), opts.UserHandler.ListSets)
	r.GET("/sets/:id", authMiddleware.IsAuthenticated(), opts.UserHandler.ListCards)

	return r

}
