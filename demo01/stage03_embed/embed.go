package stage03embed

import (
	"context"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
)

func EmbedTtext() {
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
	texts := []string{
		"hahaha",
		"玩原神玩的",
	}
	embeddings, err := embedder.EmbedStrings(ctx, texts)
	if err != nil {
		panic(err)
	}
	for i, embedding := range embeddings {
		println("文本", i+1, "的向量维度:", len(embedding))
	}
}
