package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
)

type database interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Store interface {
	Atomic(ctx context.Context, fn func(Store) (any, error)) (any, error)
	FlashcardSet() FlashcardSetRepository
	Flashcard() FlashcardRepository
	Session() SessionRepository
	SessionFlashcard() SessionFlashcardRepository
	FlashcardProgress() FlashcardProgressRepository
	User() UserRepository
}

type store struct {
	conn *sql.DB
	db   database
}

func NewStore(db *sql.DB) Store {
	return &store{
		conn: db,
		db:   db,
	}
}

func (s *store) Atomic(ctx context.Context, fn func(Store) (any, error)) (any, error) {
	tx, err := s.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, apperror.ErrInternalServerError
	}

	defer commitOrRollback(tx, recover(), &err)

	result, err := fn(&store{conn: s.conn, db: tx})

	return result, err
}

func commitOrRollback(tx *sql.Tx, p interface{}, err *error) {
	if p != nil {
		tx.Rollback()
		panic(p)
	}

	if *err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			*err = fmt.Errorf("tx err: %v, rollback err: %v", err, rollbackErr)
		}
		return
	}

	*err = tx.Commit()
}

func (s *store) FlashcardSet() FlashcardSetRepository {
	return NewFlashcardSetRepository(s.db)
}

func (s *store) Flashcard() FlashcardRepository {
	return NewFlashcardRepository(s.db)
}

func (s *store) Session() SessionRepository {
	return NewSessionRepository(s.db)
}

func (s *store) SessionFlashcard() SessionFlashcardRepository {
	return NewSessionFlashcardRepository(s.db)
}

func (s *store) FlashcardProgress() FlashcardProgressRepository {
	return NewFlashcardProgressRepository(s.db)
}

func (s *store) User() UserRepository {
	return NewUserRepository(s.db)
}
