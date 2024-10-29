package usecase

import (
	"context"
	"errors"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/dto"
	"github.com/NikenCarolina/flashcard-be/internal/model"
	"github.com/NikenCarolina/flashcard-be/internal/util"
)

type UserAuthUseCase interface {
	Register(ctx context.Context, req dto.User) error
	Login(ctx context.Context, req dto.User) (*string, error)
	Profile(ctx context.Context, userID int) (*dto.Profile, error)
}

func (u *userUseCase) Register(ctx context.Context, req dto.User) error {
	err := util.IsPasswordValid(req.Password)
	if err != nil {
		return err
	}

	userRepo := u.store.User()
	_, err = userRepo.FindByUsername(ctx, req.Username)
	if errors.Is(err, apperror.ErrInternalServerError) {
		return err
	}

	if err == nil {
		return apperror.ErrUserHasRegister
	}

	var user model.User
	user.Username = req.Username
	user.PasswordHash, err = u.bycryptProvider.Hash(req.Password)
	if err != nil {
		return err
	}
	createdUser, err := userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	set, err := u.CreateSet(ctx, createdUser.UserID)
	if err == nil {
		_, err := u.CreateCard(ctx, createdUser.UserID, set.FlashcardSetID)
		if err == nil {
			return nil
		}
	}

	return nil
}

func (u *userUseCase) Login(ctx context.Context, req dto.User) (*string, error) {
	userRepo := u.store.User()
	user, err := userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, apperror.ErrInternalServerError) {
			return nil, err
		}
		return nil, apperror.ErrInvalidUsernamePassword
	}

	err = u.bycryptProvider.CompareHashAndPassword(user.PasswordHash, req.Password)
	if err != nil {
		return nil, apperror.ErrInvalidUsernamePassword
	}

	jwtToken, err := u.jwtProvider.Sign(int64(user.UserID))
	if err != nil {
		return nil, err
	}

	return &jwtToken, nil
}

func (u *userUseCase) Profile(ctx context.Context, userID int) (*dto.Profile, error) {
	userRepo := u.store.User()
	user, err := userRepo.FindById(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &dto.Profile{Username: user.Username}, nil
}
