package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	FindById(ctx context.Context, userID int) (*model.User, error)
}

type userRepository struct {
	db database
}

func NewUserRepository(db database) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user model.User) (*model.User, error) {
	query := `
		INSERT INTO
			"users"
			("username", "password_hash")
		VALUES
			($1, $2)	
		RETURNING
			"user_id"
	`

	if err := r.db.QueryRowContext(ctx, query, user.Username, user.PasswordHash).Scan(&user.UserID); err != nil {
		return nil, apperror.ErrInternalServerError
	}

	return &user, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	query := `
		SELECT
		"user_id",
			"username",
			"password_hash"
		FROM
			"users"
		WHERE
			"username" = $1
	`

	var user model.User
	if err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.UserID,
		&user.Username,
		&user.PasswordHash,
	); err != nil {
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, apperror.ErrInternalServerError
	}

	return &user, nil
}

func (r *userRepository) FindById(ctx context.Context, userID int) (*model.User, error) {
	query := `
		SELECT
		"user_id",
			"username",
			"password_hash"
		FROM
			"users"
		WHERE
			"user_id" = $1
	`

	var user model.User
	if err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&user.UserID,
		&user.Username,
		&user.PasswordHash,
	); err != nil {
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrInvalidUsernamePassword
		}
		return nil, apperror.ErrInternalServerError
	}

	return &user, nil
}
