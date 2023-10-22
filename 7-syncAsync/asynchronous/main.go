package main

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"time"
)

const (
	TypePrintMessage = "print:message"
)

func NewPrintMessageTask(msg string) *asynq.Task {
	payload := map[string]interface{}{
		"message": msg,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	return asynq.NewTask(TypePrintMessage, payloadBytes)
}

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

	t1 := NewPrintMessageTask("Message 1")
	t2 := NewPrintMessageTask("Message 2")
	t3 := NewPrintMessageTask("Message 3")

	client.Enqueue(t1)
	time.Sleep(1 * time.Second)
	client.Enqueue(t2)
	time.Sleep(1 * time.Second)
	client.Enqueue(t3)
}
