package main

import (
	"github.com/bitly/go-nsq"
	"log"
	"encoding/json"
	"os"
)

type Task struct {
	Url string `json:"url"`
	Target string `json:"target"`
	Path string `json:"path"`
	Device string `json:"device"`
}

func handle(c *Config) {
	r, err := nsq.NewConsumer("topic", "channel", nsq.NewConfig())
	if err != nil {
		log.Fatal(err.Error())
	}

	browser := NewMacBrowser()	
	storage := NewStorage(c)
	
	r.AddConcurrentHandlers(nsq.HandlerFunc(func (m *nsq.Message) error {
		task := &Task{}
		json.Unmarshal(m.Body, &task)
		log.Println(task)
		browser.Snapshot(task.Url, task.Target, c.Timeout)
		if _, err := os.Stat(task.Target); task.Path != "" && err == nil {
			// storage put
			storage.put(task.Path, task.Target)
		}
		log.Println("task done.");
		return nil
	}), c.Workers)

	err1 := r.ConnectToNSQD(c.NsqAddr)
	
	if err1 != nil {
		log.Fatal(err1.Error())
	}

	<-r.StopChan
}