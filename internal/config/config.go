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
}

type AppConfig struct {
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
	}
}

func InitAppConfig() *AppConfig {
	gracePeriod, err := time.ParseDuration(os.Getenv("SERVER_GRACE_PERIOD"))
	if err != nil {
		log.Fatal("failed to parse SERVER_GRACE_PERIOD")
	}
	return &AppConfig{
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
	return &cors.Config{
		AllowOrigins: []string{os.Getenv("FRONTEND_URL")},
		AllowMethods: strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), " "),
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
