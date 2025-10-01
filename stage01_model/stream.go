package stage01model

import (
	"context"
	"demo01/env"
	"fmt"

	"github.com/cloudwego/eino/schema"
)

func Stream() {

	model, err := env.EnvNewChatModel()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	reader, err := model.Stream(ctx, []*schema.Message{
		schema.SystemMessage("你是一个专业的恋爱聊天机器人"),
		schema.UserMessage("你好,你是谁"),
	})
	if err != nil {
		panic(err)
	}
	defer reader.Close() // 注意要关闭
	for {
		chunk, err := reader.Recv()
		if err != nil {
			break
		}
		if chunk.Content != "" {
			fmt.Print(chunk.Content)
		}
		if usage := chunk.ResponseMeta.Usage; usage != nil {
			println("提示 Tokens:", usage.PromptTokens)
			println("生成 Tokens:", usage.CompletionTokens)
			println("总 Tokens:", usage.TotalTokens)
		}
	}

}
