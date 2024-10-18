package handler

import (
	"database/sql"

	"github.com/NikenCarolina/flashcard-be/internal/config"
	"github.com/NikenCarolina/flashcard-be/internal/repository"
	"github.com/NikenCarolina/flashcard-be/internal/usecase"
)

type HandlerOpts struct {
	*UserHandler
}

func Init(db *sql.DB, config *config.Config) *HandlerOpts {
	store := repository.NewStore(db)
	userUseCase := usecase.NewUserUseCase(store)
	return &HandlerOpts{
		UserHandler: NewUserHandler(userUseCase),
	}
}
