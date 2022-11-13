package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DelusionalOptimist/oura/internal/queue"
	s "github.com/DelusionalOptimist/oura/internal/store"
)

func RunServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/new", NewQueue)
	mux.HandleFunc("/push", PushToQueue)
	mux.HandleFunc("/pull", PullFromQueue)

	log.Println("Starting server on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return err
	}

	return nil
}

// METHOD: GET
// creates a new queue
func NewQueue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	topic := r.Form.Get("topic")
	if topic == "" {
		http.Error(w, "Queue topic cannot be empty", http.StatusBadRequest)
		return
	}

	q := queue.NewQueue()
	s.QueueStore[topic] = q
}

// METHOD: POST
// Pushes a message to the queue
func PushToQueue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	msg := &queue.Message{}
	err := json.NewDecoder(r.Body).Decode(msg)
	if err != nil {
		http.Error(w, "Unable to unmarshal message", http.StatusBadRequest)
		return
	}

	s.QueueStore[msg.Topic].QueueEnqueue(*msg)
	return
}

// METHOD: GET
// Retrieves the latest message from the queue
func PullFromQueue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	topic := r.Form.Get("topic")
	if topic == "" {
		http.Error(w, "Queue topic cannot be empty", http.StatusBadRequest)
		return
	}

	msg, err := s.QueueStore[topic].QueueDeque()
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	} else if err.Error() == "Queue empty" {
		http.Error(w, "Queue empty", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(msg)
	return
}
