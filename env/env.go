package env

import (
	"context"
	embed "github.com/cloudwego/eino-ext/components/embedding/ark"
	cmodel "github.com/cloudwego/eino-ext/components/model/ark"
	"os"
	"time"
)

func EnvNewChatModel() (*cmodel.ChatModel, error) {
	ctx := context.Background()
	tout := time.Second * 30
	chatmodel, err := cmodel.NewChatModel(ctx, &cmodel.ChatModelConfig{
		Timeout: &tout,
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   "doubao-seed-1-6-flash-250828",
	})
	return chatmodel, err
}
func EnvNewEmbed() (*embed.Embedder, error) {
	ctx := context.Background()
	tout := time.Second * 30
	embedder, err := embed.NewEmbedder(ctx, &embed.EmbeddingConfig{
		Timeout: &tout,
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   "doubao-embedding-text-240715",
	})
	return embedder, err
}
