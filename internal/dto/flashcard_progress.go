package dto

import "time"

type FlashcardProgress struct {
	RepetitionNumber *int64     `json:"repetition_number" binding:"required,number"`
	EasinessFactor   float64    `json:"easiness_factor" binding:"required"`
	Interval         *int64     `json:"interval" binding:"required,number"`
	LastReview       *time.Time `json:"last_review"`
	DueDate          *time.Time `json:"due_date"`
}
