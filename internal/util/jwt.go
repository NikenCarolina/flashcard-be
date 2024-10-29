package util

import (
	"time"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	UserID int64
	jwt.RegisteredClaims
}

type JwtProvider interface {
	Sign(userID int64) (string, error)
	Parse(tokenString string) (*jwtClaims, error)
}

type jwtProvider struct {
	config config.JwtConfig
}

func NewJwtProvider(jwtConfig config.JwtConfig) JwtProvider {
	return &jwtProvider{
		config: jwtConfig,
	}
}

func (p *jwtProvider) Sign(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    p.config.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(p.config.ExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	s, err := token.SignedString([]byte(p.config.SecretKey))
	if err != nil {
		return "", err
	}

	return s, nil
}

func (p *jwtProvider) Parse(tokenString string) (*jwtClaims, error) {
	parser := jwt.NewParser(
		jwt.WithValidMethods(p.config.AllowedAlgs),
		jwt.WithIssuer(p.config.Issuer),
		jwt.WithIssuedAt(),
		jwt.WithExpirationRequired(),
	)

	token, err := parser.ParseWithClaims(tokenString, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(p.config.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, apperror.ErrInvalidToken
}
