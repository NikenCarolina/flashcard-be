package repository

type CardSetRepository interface{}
type cardSetRepository struct{}

func NewCardSetRepository() *cardSetRepository {
	return &cardSetRepository{}
}
