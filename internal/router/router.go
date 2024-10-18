package router

import (
	"net/http"

	"github.com/NikenCarolina/flashcard-be/internal/handler"
	"github.com/gin-gonic/gin"
)

func Init(opts *handler.HandlerOpts) http.Handler {
	r := gin.New()
	r.GET("/sets", opts.UserHandler.ListSets)
	return r

}
