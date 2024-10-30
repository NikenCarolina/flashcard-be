package dto

type FlashcardUri struct {
	FlashcardSetID int `uri:"set_id" binding:"required"`
	FlashcardID    int `uri:"card_id" binding:"required"`
}

type Flashcard struct {
	FlashcardID    int     `json:"id" binding:"required"`
	FlashcardSetID int     `json:"flashcard_set_id" binding:"required"`
	Term           *string `json:"term" binding:"required"`
	Definition     *string `json:"definition" binding:"required"`
}

type FlashcardUpdateResponse struct {
	Status      int64 `json:"status"`
	FlashcardID int   `json:"id"`
}
