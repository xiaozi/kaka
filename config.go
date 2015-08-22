package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	AccessKey string
	SecretKey string
	Bucket    string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{AccessKey: os.Getenv("ACCESS_KEY"), SecretKey: os.Getenv("SECRET_KEY"), Bucket: os.Getenv("BUCKET")}
}
