package dto

type FlashcardSetUri struct {
	FlashcardSetID int `uri:"set_id" binding:"required"`
}

type FlashcardSetRequest struct {
	FlashcardSetID int `json:"flashcard_set_id" binding:"required"`
}

type FlashcardSet struct {
	FlashcardSetID int    `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
}
