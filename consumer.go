package main

import (
	"github.com/bitly/go-nsq"
	"log"
	"encoding/json"
)

type Task struct {
	Url string `json:"url"`
	Path string `json:"path"`
	Device string `json:"device"`
}

func handle(c *Config) {
	r, err := nsq.NewConsumer("topic", "channel", nsq.NewConfig())
	if err != nil {
		log.Fatal(err.Error())
	}

	browser := NewMacBrowser()	
	
	r.AddConcurrentHandlers(nsq.HandlerFunc(func (m *nsq.Message) error {
		task := &Task{}
		json.Unmarshal(m.Body, &task)
		log.Println(task)
		browser.Snapshot(task.Url, task.Path)
		return nil
	}), c.Workers)

	err1 := r.ConnectToNSQD(c.NsqAddr)
	
	if err1 != nil {
		log.Fatal(err1.Error())
	}

	<-r.StopChan
}