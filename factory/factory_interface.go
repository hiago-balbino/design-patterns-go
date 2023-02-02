package main

import "fmt"

// MessageType is a constant for all types of messages.
type MessageType int

const (
	EMAIL MessageType = iota
	TEXT
)

// Message is an interface to send a message.
type Message interface {
	// Send is a function to be implemented for all message types.
	Send()
}

type emailMessage struct{}

// Send is a function implemented for the email message type.
func (e emailMessage) Send() {
	fmt.Println("sending email message")
}

type textMessage struct{}

// Send is a function implemented for the text message type.
func (t textMessage) Send() {
	fmt.Println("sending text message")
}

// NewMessage is a constructor/factory to create a new instance of implemented Message interface given the MessageType.
func NewMessage(messageType MessageType) Message {
	if messageType == TEXT {
		return textMessage{}
	}
	return emailMessage{}
}

func main() {
	textMessage := NewMessage(TEXT)
	textMessage.Send()

	emailMessage := NewMessage(EMAIL)
	emailMessage.Send()
}
