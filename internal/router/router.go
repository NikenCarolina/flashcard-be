package router

import (
	"net/http"

	"github.com/NikenCarolina/flashcard-be/internal/config"
	"github.com/NikenCarolina/flashcard-be/internal/handler"
	"github.com/NikenCarolina/flashcard-be/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(opts *handler.HandlerOpts, config *config.Config) http.Handler {
	r := gin.Default()
	r.ContextWithFallback = true

	middlewares := []gin.HandlerFunc{
		cors.New(*config.Cors),
		middleware.Error(),
	}
	r.Use(middlewares...)

	authMiddleware := middleware.NewAuthMiddleware()

	r.GET("/sets", authMiddleware.IsAuthenticated(), opts.UserHandler.ListSets)
	r.GET("/sets/:id", authMiddleware.IsAuthenticated(), opts.UserHandler.GetSetById)
	r.GET("/sets/:id/cards", authMiddleware.IsAuthenticated(), opts.UserHandler.ListCards)

	return r
}
