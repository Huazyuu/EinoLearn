package stage02template

import (
	"context"
	"demo01/env"
	"fmt"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func TemplateChatModel() {
	ctx := context.Background()
	tmp := prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}"),
		schema.MessagesPlaceholder("history_key", false),
		&schema.Message{
			Role:    schema.User,
			Content: "请帮帮我，{role},{task}",
		},
	)
	params := map[string]any{
		"role": "古诗词高手",
		"task": "写一首诗你擅长的诗词",
		"history_key": []*schema.Message{
			{Role: schema.User, Content: "你擅长什么类型的诗词?"},
			{Role: schema.Assistant, Content: "我擅长唐诗"},
		},
	}
	messages, err := tmp.Format(ctx, params)
	if err != nil {
		panic(err)
	}

	model, err := env.EnvNewChatModel()
	if err != nil {
		panic(err)
	}
	answer, err := model.Generate(ctx, messages)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer.Content)
	if answer.ResponseMeta.Usage != nil {
		println("提示 Tokens:", answer.ResponseMeta.Usage.PromptTokens)
		println("生成 Tokens:", answer.ResponseMeta.Usage.CompletionTokens)
		println("总 Tokens:", answer.ResponseMeta.Usage.TotalTokens)
	}
}
