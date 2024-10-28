package model

import "github.com/NikenCarolina/flashcard-be/internal/dto"

type FlashcardSet struct {
	FlashcardSetID int
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

func (m *FlashcardSet) LoadFromDto(s dto.FlashcardSet) {
	m.FlashcardSetID = s.FlashcardSetID
	m.Title = s.Title
	m.Description = s.Description
}
