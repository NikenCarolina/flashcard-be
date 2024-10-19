package dto

type FlashcardSet struct {
	FlashcardSetID int64  `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
}
