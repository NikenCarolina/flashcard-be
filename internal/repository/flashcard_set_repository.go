package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/model"
)

type FlashcardSetRepository interface {
	GetAll(ctx context.Context, userID int) ([]model.FlashcardSet, error)
	GetById(ctx context.Context, userID int, setID int) (*model.FlashcardSet, error)
	CheckExists(ctx context.Context, userID, setID int) (bool, error)
}

type flashcardSetRepository struct {
	db database
}

func NewFlashcardSetRepository(db database) FlashcardSetRepository {
	return &flashcardSetRepository{db}
}

func (r *flashcardSetRepository) GetAll(ctx context.Context, userID int) ([]model.FlashcardSet, error) {
	flashcardSets := []model.FlashcardSet{}
	query := `
		SELECT 
			"flashcard_set_id", 
			"title", 
			"description" 
		FROM "flashcard_sets" 
		WHERE "user_id" = $1`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, apperror.ErrInternalServerError
	}
	defer rows.Close()

	for rows.Next() {
		var flashcardSet model.FlashcardSet
		if err := rows.Scan(
			&flashcardSet.FlashcardSetID,
			&flashcardSet.Title,
			&flashcardSet.Description,
		); err != nil {
			return nil, apperror.ErrInternalServerError
		}
		flashcardSets = append(flashcardSets, flashcardSet)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.ErrInternalServerError
	}

	return flashcardSets, nil
}

func (r *flashcardSetRepository) CheckExists(ctx context.Context, userID, setID int) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM flashcard_sets
			WHERE user_id = $1 AND flashcard_set_id = $2
		)
	`

	var exists bool
	if err := r.db.QueryRowContext(ctx, query, userID, setID).Scan(&exists); err != nil {
		return false, apperror.ErrInternalServerError
	}

	return exists, nil
}

func (r *flashcardSetRepository) GetById(ctx context.Context, userID, setID int) (*model.FlashcardSet, error) {
	query := `
		SELECT 
			"flashcard_set_id", 
			"title", 
			"description" 
		FROM "flashcard_sets"
		WHERE "user_id" = $1 AND "flashcard_set_id" = $2
	`
	flashcardSet := &model.FlashcardSet{}
	dest := []interface{}{&flashcardSet.FlashcardSetID, &flashcardSet.Title, &flashcardSet.Description}

	if err := r.db.QueryRowContext(ctx, query, userID, setID).Scan(dest...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrNotFound
		}
		return nil, apperror.ErrInternalServerError
	}

	return flashcardSet, nil
}
