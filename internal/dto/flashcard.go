package dto

type FlashcardUri struct {
	FlashcardSetID int64 `uri:"set_id" binding:"required"`
	FlashcardID    int64 `uri:"card_id" binding:"required"`
}

type Flashcard struct {
	FlashcardID    int64  `json:"id"`
	FlashcardSetID int64  `json:"flashcard_set_id"`
	Term           string `json:"term"`
	Definition     string `json:"definition"`
}

type FlashcardUpdateResponse struct {
	Status      int64 `json:"status"`
	FlashcardID int64 `json:"id"`
}
