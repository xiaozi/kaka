package main

import (
	"golang.org/x/net/context"
	"log"
	"qiniupkg.com/api.v7/kodo"
)

type Storage struct {
	Bucket kodo.Bucket
}

func NewStorage(c *Config) *Storage {
	kodo.SetMac(c.AccessKey, c.SecretKey)

	zone := 0
	client := kodo.New(zone, nil)
	bucket := client.Bucket(c.Bucket)

	return &Storage{Bucket: bucket}
}

func (s *Storage) put(remotePath string, localPath string) error {
	ctx := context.Background()
	err := s.Bucket.PutFile(ctx, nil, remotePath, localPath, nil)
	if err != nil {
		log.Print("upload failed:", err)
	}
	return err
}
