package main 

import (
	"time"
)

// Message doc ...
type Message struct{
	ID int `josn:"id"`
	Author string	`json:"content"`
	Content string 	`json:"content"`
	Timestap time.Time `json:"timestap"`
}

func newMessage(id int, author string, content string) Message {
	return Message{
		ID: id, 
		Author: author, 
		Content: content, 
		Timestap: time.Now(), 
	}
}