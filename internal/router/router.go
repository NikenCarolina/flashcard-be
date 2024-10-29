package router

import (
	"net/http"

	"github.com/NikenCarolina/flashcard-be/internal/config"
	"github.com/NikenCarolina/flashcard-be/internal/handler"
	"github.com/NikenCarolina/flashcard-be/internal/middleware"
	"github.com/NikenCarolina/flashcard-be/internal/util"
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

	r.GET("/auth/logout", opts.UserHandler.Logout)

	r.POST("/auth/login", opts.UserHandler.Login)
	r.POST("/auth/register", opts.UserHandler.Register)

	jwtProvider := util.NewJwtProvider(*config.Jwt)
	authMiddleware := middleware.NewAuthMiddleware(jwtProvider)

	r.DELETE("/sets/:set_id", authMiddleware.IsAuthenticated(), opts.UserHandler.DeleteSet)
	r.DELETE("/sets/:set_id/cards/:card_id", authMiddleware.IsAuthenticated(), opts.UserHandler.DeleteCard)

	r.GET("/profile", authMiddleware.IsAuthenticated(), opts.UserHandler.Profile)
	r.GET("/sets", authMiddleware.IsAuthenticated(), opts.UserHandler.ListSets)
	r.GET("/sets/:set_id", authMiddleware.IsAuthenticated(), opts.UserHandler.GetSetById)
	r.GET("/sets/:set_id/cards", authMiddleware.IsAuthenticated(), opts.UserHandler.ListCards)

	r.POST("/sessions", authMiddleware.IsAuthenticated(), opts.UserHandler.StartSession)
	r.POST("/sets", authMiddleware.IsAuthenticated(), opts.UserHandler.CreateSet)
	r.POST("/sets/:set_id/cards", authMiddleware.IsAuthenticated(), opts.UserHandler.CreateCard)

	r.PUT("/sessions/:session_id", authMiddleware.IsAuthenticated(), opts.UserHandler.EndSession)
	r.PUT("/sets/:set_id", authMiddleware.IsAuthenticated(), opts.UserHandler.UpdateSet)
	r.PUT("/sets/:set_id/cards", authMiddleware.IsAuthenticated(), opts.UserHandler.BulkUpdateCard)
	r.PUT("/sets/:set_id/cards/:card_id", authMiddleware.IsAuthenticated(), opts.UserHandler.UpdateCard)

	return r
}
