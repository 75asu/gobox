package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"time"
)

const (
	TypePrintMessage = "print:message"
)

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypePrintMessage, func(ctx context.Context, task *asynq.Task) error {
		var payload map[string]interface{}
		if err := json.Unmarshal(task.Payload(), &payload); err != nil {
			return fmt.Errorf("could not unmarshal payload: %v", err)
		}
		msg, ok := payload["message"].(string)
		if !ok {
			return fmt.Errorf("cannot get string 'message' from payload")
		}

		time.Sleep(1 * time.Second)
		fmt.Println(msg)

		return nil 
	})

	if err := srv.Run(mux); err != nil {
		fmt.Printf("could not run server: %v", err)
	}
}
