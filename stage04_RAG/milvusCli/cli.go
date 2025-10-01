package milvusCli

import (
	"context"
	"log"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

var MilvusCli client.Client

func init() {
	ctx := context.Background()
	cli, err := client.NewClient(ctx, client.Config{
		Address: "127.0.0.1:19530",
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	MilvusCli = cli
}
