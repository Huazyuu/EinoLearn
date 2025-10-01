package env

import (
	"context"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ark"
)

func EnvNewChatModel() (*ark.ChatModel, error) {
	ctx := context.Background()
	tout := time.Second * 30
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		Timeout: &tout,
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   "doubao-seed-1-6-flash-250828",
	})
	return model, err
}
