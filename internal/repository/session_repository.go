package repository

import (
	"context"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/model"
)

type SessionRepository interface {
	CheckExists(ctx context.Context, sessionID, userID, setID int) (bool, error)
	Create(ctx context.Context, userID, setID int) (*model.Session, error)
	EndBySetId(ctx context.Context, setID int) error
	EndById(ctx context.Context, sessionID int) error
}

type sessionRepository struct {
	db database
}

func NewSessionRepository(db database) SessionRepository {
	return &sessionRepository{db}
}

func (r *sessionRepository) CheckExists(ctx context.Context, sessionID, userID, setID int) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM "sessions" 
			WHERE "session_id" = $1 AND "user_id" = $2 AND "flashcard_set_id" = $3
		)
	`

	var exists bool
	if err := r.db.QueryRowContext(ctx, query, sessionID, userID, setID).Scan(&exists); err != nil {
		return false, apperror.ErrInternalServerError
	}

	return exists, nil
}

func (r *sessionRepository) Create(ctx context.Context, userID, setID int) (*model.Session, error) {
	query := `
		INSERT INTO "sessions"
			("user_id", "flashcard_set_id")
		VALUES
			($1, $2)
		RETURNING
			"session_id"
	`

	var res model.Session
	if err := r.db.QueryRowContext(ctx, query, userID, setID).Scan(&res.SessionID); err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *sessionRepository) EndBySetId(ctx context.Context, setID int) error {
	query := `
		UPDATE "sessions" 
		SET "end_at" = NOW()
		WHERE "flashcard_set_id" = $1 AND "end_at" ISNULL 
	`
	if _, err := r.db.ExecContext(ctx, query, setID); err != nil {
		return apperror.ErrInternalServerError
	}
	return nil
}

func (r *sessionRepository) EndById(ctx context.Context, sessionID int) error {
	query := `
		UPDATE "sessions" 
		SET "end_at" = NOW()
		WHERE "session_id" = $1 AND "end_at" ISNULL 
	`
	res, err := r.db.ExecContext(ctx, query, sessionID)
	if err != nil {
		return apperror.ErrInternalServerError
	}
	rowNum, err := res.RowsAffected()
	if err != nil {
		return apperror.ErrInternalServerError
	}
	if rowNum == 0 {
		return apperror.ErrNotFound
	}

	return nil
}
