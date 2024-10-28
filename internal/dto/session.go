package dto

type Session struct {
	SessionID  int                `json:"session_id"`
	Flashcards []SessionFlashcard `json:"flashcards"`
}

type SessionFlashcard struct {
	Flashcard
	FlashcardProgress
}

type EndSessionFlashcard struct {
	SessionFlashcard
	IsCorrect *bool `json:"is_correct" binding:"required"`
}

type SessionUri struct {
	SessionID int `uri:"session_id" binding:"required"`
}

type EndSessionRequest struct {
	SetID      int                   `json:"flashcard_set_id" binding:"required"`
	Flashcards []EndSessionFlashcard `json:"flashcards" binding:"dive,required"`
}
