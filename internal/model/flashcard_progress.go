package model

import (
	"time"

	"github.com/NikenCarolina/flashcard-be/internal/dto"
)

type FlashcardProgress struct {
	FlashcardSetID   int64
	FlashcardID      int64
	RepetitionNumber int64
	EasinessFactor   float64
	Interval         int64
	LastReview       *time.Time
	DueDate          *time.Time
}

func (f *FlashcardProgress) ToDto() *dto.FlashcardProgress {
	return &dto.FlashcardProgress{
		RepetitionNumber: &f.RepetitionNumber,
		EasinessFactor:   f.EasinessFactor,
		Interval:         &f.Interval,
		LastReview:       f.LastReview,
		DueDate:          f.DueDate,
	}
}
