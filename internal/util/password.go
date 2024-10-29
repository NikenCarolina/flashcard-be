package util

import (
	"regexp"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
)

var (
	hasLower   = regexp.MustCompile(`[a-z]`).MatchString
	hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString
	hasDigit   = regexp.MustCompile(`[0-9]`).MatchString
	hasSpecial = regexp.MustCompile(`[\W]`).MatchString
)

func IsPasswordValid(password string) error {
	if !hasLower(password) || !hasUpper(password) || !hasDigit(password) || !hasSpecial(password) {
		return apperror.ErrInvalidPasswordFormat
	}
	return nil
}
