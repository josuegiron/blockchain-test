package main 

import (
	"time"
)

// Message doc ...
type Message struct{
	Author string	`json:"author"`
	Content string 	`json:"content"`
	Timestap string `json:"timestap"`
}

func newMessage(author string, content string) Message {
	return Message{
		Author: author, 
		Content: content, 
		Timestap: time.Now().Format(time.RFC3339), 
	}
}

// Queue doc ...
type Queue struct {
	Messages []Message `json:"messages"`
}

func (queue *Queue)clean() {
	queue.Messages = nil
}