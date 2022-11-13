package queue

import "fmt"

// a message stored in the queue
type Message struct {
	// Topic of the queue that this message belongs to
	Topic string `json:"topic"`

	// ID of the message. We'll use it later (Prolly)
	ID string `json:"id"`

	// Content stores the message
	// TODO: encryption?
	Content string `json:"content"`
}

// this is our Queue
type Queue struct {
	// stores the messages
	Messages []Message
}

func NewQueue() *Queue {
	return &Queue{
		Messages: make([]Message, 0),
	}
}

// basic queue methods
func (q *Queue) QueueDeque() (Message, error) {
	if len(q.Messages) == 0 {
		return Message{}, fmt.Errorf("Queue empty")
	}
	msg := q.Messages[0]
	q.Messages = q.Messages[1:]
	return msg, nil
}

func (q *Queue) QueueEnqueue(msg Message) {
	q.Messages = append(q.Messages, msg)
}
