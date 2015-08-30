package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	AccessKey string
	SecretKey string
	Bucket    string
	NsqAddr   string
	Workers   int
	Timeout   int
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	workers, _ := strconv.ParseInt(os.Getenv("WORKERS"), 10, 0)
	timeout, _ := strconv.ParseInt(os.Getenv("TIMEOUT"), 10, 0)
	return &Config{
		AccessKey: os.Getenv("ACCESS_KEY"),
		SecretKey: os.Getenv("SECRET_KEY"),
		Bucket: os.Getenv("BUCKET"),
		NsqAddr: os.Getenv("NSQ_ADDR"),
		Workers: int(workers),
		Timeout: int(timeout),
	}
}
