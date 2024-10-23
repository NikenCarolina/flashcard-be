package dto

type FlashcardSetUri struct {
	FlashcardSetID int `uri:"id" binding:"required"`
}

type FlashcardSet struct {
	FlashcardSetID int    `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
}
