package handler

import (
	"database/sql"

	"github.com/NikenCarolina/flashcard-be/internal/config"
	"github.com/NikenCarolina/flashcard-be/internal/repository"
	"github.com/NikenCarolina/flashcard-be/internal/usecase"
	"github.com/NikenCarolina/flashcard-be/internal/util"
)

type HandlerOpts struct {
	*UserHandler
}

func Init(db *sql.DB, config *config.Config) *HandlerOpts {
	store := repository.NewStore(db)
	jwtProvider := util.NewJwtProvider(*config.Jwt)
	bycryptProvider := util.NewBcryptProvider(config.Bycrypt.Cost)
	userUseCase := usecase.NewUserUseCase(store, *config.Flashcard, jwtProvider, bycryptProvider)
	return &HandlerOpts{
		UserHandler: NewUserHandler(userUseCase, config),
	}
}
