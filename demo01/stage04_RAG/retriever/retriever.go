package retriever

import (
	"context"
	"einolearn/demo01/stage04_RAG/indexer"
	"einolearn/demo01/stage04_RAG/milvusCli"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/retriever/milvus"
	"github.com/cloudwego/eino/schema"
)

func RetrieveRAG(query string, cnt int) []*schema.Document {
	ctx := context.Background()
	timeout := time.Second * 30
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		Timeout: &timeout,
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   "doubao-embedding-text-240715",
	})
	if err != nil {
		panic(err)
	}
	retriever, err := milvus.NewRetriever(ctx, &milvus.RetrieverConfig{
		Client:      milvusCli.MilvusCli,
		Collection:  indexer.Collection,
		Partition:   nil,
		VectorField: "vector",
		OutputFields: []string{
			"id",
			"content",
			"metadata",
		},
		TopK:      cnt,
		Embedding: embedder,
	})
	if err != nil {
		panic(err)
	}
	res, err := retriever.Retrieve(ctx, query)
	if err != nil {
		panic(err)
	}
	return res

}
