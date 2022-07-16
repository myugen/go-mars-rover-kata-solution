package greetings

import "fmt"

const messageTemplate = "Hello %s!"

type Message struct {
	to string
}

func (g Message) Say() string {
	to := "everyone"
	if g.to != "" {
		to = g.to
	}

	return fmt.Sprintf(messageTemplate, to)
}

func NewMessage(to string) Message {
	return Message{to: to}
}
