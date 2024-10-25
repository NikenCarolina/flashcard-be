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
	r.GET("/sets/:set_id", authMiddleware.IsAuthenticated(), opts.UserHandler.GetSetById)
	r.GET("/sets/:set_id/cards", authMiddleware.IsAuthenticated(), opts.UserHandler.ListCards)
	r.POST("/sets/:set_id/cards", authMiddleware.IsAuthenticated(), opts.UserHandler.CreateCard)
	r.PUT("/sets/:set_id/cards/:card_id", authMiddleware.IsAuthenticated(), opts.UserHandler.UpdateCard)
	r.PUT("/sets/:set_id/cards", authMiddleware.IsAuthenticated(), opts.UserHandler.BulkUpdateCard)
	r.DELETE("/sets/:set_id/cards/:card_id", authMiddleware.IsAuthenticated(), opts.UserHandler.DeleteCard)

	return r
}
