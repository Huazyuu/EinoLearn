package stage06

import (
	"context"
	"einolearn/demo01/env"
	"fmt"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

func Chain() {
	chatmodel, err := env.EnvNewChatModel()
	if err != nil {
		panic(err)
	}
	lambda := compose.InvokableLambda(func(ctx context.Context, input string) (output []*schema.Message, err error) {
		desuwa := input + `回答结尾加上"desu"`
		output = []*schema.Message{
			{
				Role:    schema.User,
				Content: desuwa,
			},
		}
		return output, nil
	})
	chain := compose.NewChain[string, *schema.Message]()
	chain.AppendLambda(lambda).AppendChatModel(chatmodel)
	r, err := chain.Compile(context.Background())
	if err != nil {
		panic(err)
	}
	ans, err := r.Invoke(context.Background(), "你好,你的名字是什么")
	if err != nil {
		panic(err)
	}
	fmt.Println(ans.Content)
}
