package dto

type Flashcard struct {
	FlashcardID    int64  `json:"id"`
	FlashcardSetID int64  `json:"flashcard_set_id"`
	Term           string `json:"term"`
	Definition     string `json:"definition"`
}
