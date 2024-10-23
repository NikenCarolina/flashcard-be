package model

import "github.com/NikenCarolina/flashcard-be/internal/dto"

type Flashcard struct {
	FlashcardID    int64
	FlashcardSetID int64
	Term           string
	Definition     string
}

func (m *Flashcard) ToDto() *dto.Flashcard {
	return &dto.Flashcard{
		FlashcardID:    m.FlashcardID,
		FlashcardSetID: m.FlashcardSetID,
		Term:           m.Term,
		Definition:     m.Definition,
	}
}
