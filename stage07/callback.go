package stage07

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/callbacks"
)

func genCallback() callbacks.Handler {
	return callbacks.NewHandlerBuilder().
		OnStartFn(func(ctx context.Context, info *callbacks.RunInfo, input callbacks.CallbackInput) context.Context {
			fmt.Printf("当前%s节点输入:%s\n", info.Component, input)
			return ctx
		}).
		OnEndFn(func(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) context.Context {
			fmt.Printf("当前%s节点输出:%s\n", info.Component, output)
			return ctx
		}).Build()
}
