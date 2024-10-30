package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

type Config struct {
	App       *AppConfig
	Database  *DatabaseConfig
	Cors      *cors.Config
	Flashcard *FlashcardConfig
	Jwt       *JwtConfig
	Bycrypt   *BycryptConfig
}

type AppConfig struct {
	AuthRedirectURL   string
	DomainName        string
	ServerAddress     string
	ServerGracePeriod time.Duration
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     uint16
}

type FlashcardConfig struct {
	RepetitionNumber int64
	EasinessFactor   float64
	Interval         int64
}

type JwtConfig struct {
	Issuer         string
	SecretKey      string
	AllowedAlgs    []string
	ExpireDuration time.Duration
}

type BycryptConfig struct {
	Cost int
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env")
	}
	return &Config{
		App:       InitAppConfig(),
		Database:  InitDatabaseConfig(),
		Cors:      InitCorsConfig(),
		Flashcard: InitFlashcardConfig(),
		Jwt:       InitJwtConfig(),
		Bycrypt:   InitBycryptConfig(),
	}
}

func InitAppConfig() *AppConfig {
	gracePeriod, err := time.ParseDuration(os.Getenv("SERVER_GRACE_PERIOD"))
	if err != nil {
		log.Fatal("failed to parse SERVER_GRACE_PERIOD")
	}
	return &AppConfig{
		DomainName:        os.Getenv("DOMAIN_NAME"),
		AuthRedirectURL:   os.Getenv("AUTH_REDIRECT_URL"),
		ServerAddress:     os.Getenv("SERVER_PORT"),
		ServerGracePeriod: gracePeriod,
	}
}

func InitDatabaseConfig() *DatabaseConfig {

	port, err := strconv.ParseUint(os.Getenv("DATABASE_PORT"), 10, 16)
	if err != nil {
		log.Fatal("failed to parse DATABASE_PORT")
	}

	return &DatabaseConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Name:     os.Getenv("DATABASE_NAME"),
		Port:     uint16(port),
	}
}

func InitCorsConfig() *cors.Config {

	allowCredentials, err := strconv.ParseBool(os.Getenv("CORS_ALLOW_CREDENTIALS"))
	if err != nil {
		log.Fatal("failed to parse CORS_ALLOW_CREDENTIALS")
	}

	return &cors.Config{
		AllowOrigins:     strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ","),
		AllowMethods:     strings.Split(os.Getenv("CORS_ALLOW_METHODS"), ","),
		AllowHeaders:     strings.Split(os.Getenv("CORS_ALLOW_HEADERS"), ","),
		AllowCredentials: allowCredentials,
	}
}

func InitFlashcardConfig() *FlashcardConfig {
	repetitionNumber, err := strconv.ParseInt(os.Getenv("REPETITION_NUMBER"), 10, 64)
	if err != nil {
		log.Fatal("failed to parse REPETITION_NUMBER")
	}

	easinessFactor, err := strconv.ParseFloat(os.Getenv("EASINESS_FACTOR"), 64)
	if err != nil {
		log.Fatal("failed to parse EASINESS_FACTOR")
	}

	interval, err := strconv.ParseInt(os.Getenv("INTERVAL"), 10, 64)
	if err != nil {
		log.Fatal("failed to parse INTERVAL")
	}

	return &FlashcardConfig{
		RepetitionNumber: repetitionNumber,
		EasinessFactor:   easinessFactor,
		Interval:         interval,
	}
}

func InitJwtConfig() *JwtConfig {
	expireDuration, err := time.ParseDuration(os.Getenv("JWT_EXPIRE_DURATION"))
	if err != nil {
		log.Fatal("failed to parse JWT_EXPIRE_DURATION")
	}

	return &JwtConfig{
		Issuer:         os.Getenv("JWT_ISSUER"),
		SecretKey:      os.Getenv("JWT_SECRET_KEY"),
		AllowedAlgs:    strings.Split(os.Getenv("JWT_ALLOWED_ALGOS"), ", "),
		ExpireDuration: expireDuration,
	}
}

func InitBycryptConfig() *BycryptConfig {
	cost, err := strconv.ParseInt(os.Getenv("BCRYPT_COST"), 10, 64)
	if err != nil {
		log.Fatal("failed to parse INTERVAL")
	}

	return &BycryptConfig{
		Cost: int(cost),
	}
}
