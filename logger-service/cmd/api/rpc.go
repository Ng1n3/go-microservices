package main

import (
	"context"
	"log"
	"time"

	"github.com/Ng1n3/go-microservices/logger-service/data"
)

type RPCServer struct{}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	ctx := context.TODO()
	_, err := collection.InsertOne(ctx, data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error writting to mongo", err)
		return err
	}

	*resp = "Processed payload via RPC successfully" + payload.Name
	return nil
}
