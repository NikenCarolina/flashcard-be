package model

import "github.com/NikenCarolina/flashcard-be/internal/dto"

type FlashcardSet struct {
	FlashcardSetID int64
	Title          string
	Description    string
}

func (m *FlashcardSet) ToDto() *dto.FlashcardSet {
	return &dto.FlashcardSet{
		FlashcardSetID: m.FlashcardSetID,
		Title:          m.Title,
		Description:    m.Description,
	}
}
