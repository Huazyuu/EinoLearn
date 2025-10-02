package stage01model

import (
	"context"
	"einolearn/demo01/env"
	"fmt"

	"github.com/cloudwego/eino/schema"
)

func ChatModel() {
	ctx := context.Background()
	model, err := env.EnvNewChatModel()
	if err != nil {
		panic(err)
	}
	msgs := []*schema.Message{
		schema.SystemMessage("你是一个专业的计算机coding机器人"),
		schema.UserMessage("你好"),
	}
	resp, err := model.Generate(ctx, msgs)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Content)
	if usage := resp.ResponseMeta.Usage; usage != nil {
		println("提示 Tokens:", usage.PromptTokens)
		println("生成 Tokens:", usage.CompletionTokens)
		println("总 Tokens:", usage.TotalTokens)
	}

}
