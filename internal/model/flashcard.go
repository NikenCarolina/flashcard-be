package model

import "github.com/NikenCarolina/flashcard-be/internal/dto"

type Flashcard struct {
	FlashcardID    int
	FlashcardSetID int
	Term           string
	Definition     string
}

func (m *Flashcard) ToDto() *dto.Flashcard {
	return &dto.Flashcard{
		FlashcardID:    m.FlashcardID,
		FlashcardSetID: m.FlashcardSetID,
		Term:           &m.Term,
		Definition:     &m.Definition,
	}
}

func (m *Flashcard) LoadFromDto(f dto.Flashcard) {
	m.FlashcardSetID = f.FlashcardSetID
	m.FlashcardID = f.FlashcardID
	m.Term = *f.Term
	m.Definition = *f.Definition
}
