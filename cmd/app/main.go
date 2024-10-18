package main

import (
	"github.com/NikenCarolina/flashcard-be/internal/config"
	"github.com/NikenCarolina/flashcard-be/internal/handler"
	"github.com/NikenCarolina/flashcard-be/internal/repository/postgres"
	"github.com/NikenCarolina/flashcard-be/internal/router"
	"github.com/NikenCarolina/flashcard-be/internal/server"
)

func main() {
	config := config.InitConfig()
	db := postgres.Init(config)
	defer db.Close()
	handler := handler.Init(db, config)
	router := router.Init(handler)
	server := server.NewServer(config, router)
	server.Run()
}
