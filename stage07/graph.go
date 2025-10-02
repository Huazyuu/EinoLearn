package stage07

import (
	"context"
	"demo01/env"
	"os"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

func Graph(ctx context.Context) *compose.Graph[map[string]string, *schema.Message] {
	g := compose.NewGraph[map[string]string, *schema.Message](compose.WithGenLocalState(genFunc))
	lambda := compose.InvokableLambda(func(ctx context.Context, input map[string]string) (map[string]string, error) {
		compose.ProcessState[*State](ctx, func(_ context.Context, state *State) error {
			state.History["tsundere_action"] = "我喜欢你"
			state.History["cute_action"] = "摸摸头"
			return nil
		})
		if input["role"] == "tsundere" {
			return map[string]string{"role": "tsundere", "content": input["content"]}, nil
		}
		if input["role"] == "cute" {
			return map[string]string{"role": "cute", "content": input["content"]}, nil
		}
		return map[string]string{"role": "user", "content": input["content"]}, nil
	})
	TsundereLambda := compose.InvokableLambda(func(ctx context.Context, input map[string]string) (output []*schema.Message, err error) {
		compose.ProcessState[*State](ctx, func(_ context.Context, state *State) error {
			input["content"] = input["content"] + state.History["tsundere_action"].(string)
			return nil
		})
		return []*schema.Message{
			{
				Role:    schema.System,
				Content: "你是一个高冷傲娇的大小姐，每次都会用傲娇的语气回答我的问题",
			},
			{
				Role:    schema.User,
				Content: input["content"],
			},
		}, nil
	})
	CuteLambda := compose.InvokableLambda(func(ctx context.Context, input map[string]string) (output []*schema.Message, err error) {
		return []*schema.Message{
			{
				Role:    schema.System,
				Content: "你是一个可爱的小女孩，每次都会用可爱的语气回答我的问题",
			},
			{
				Role:    schema.User,
				Content: input["content"],
			},
		}, nil
	})
	cutePreHandler := func(ctx context.Context, input map[string]string, state *State) (map[string]string, error) {
		input["content"] = input["content"] + state.History["cute_action"].(string)
		return input, nil
	}

	model, err := env.EnvNewChatModel()
	if err != nil {
		return nil
	}

	err = g.AddLambdaNode("lambda", lambda)
	if err != nil {
		panic(err)
	}
	err = g.AddLambdaNode("tsundere", TsundereLambda)
	if err != nil {
		panic(err)
	}
	err = g.AddLambdaNode("cute", CuteLambda, compose.WithStatePreHandler(cutePreHandler))
	if err != nil {
		panic(err)
	}
	err = g.AddChatModelNode("model", model)
	if err != nil {
		panic(err)
	}
	g.AddBranch("lambda", compose.NewGraphBranch(func(ctx context.Context, in map[string]string) (endNode string, err error) {
		if in["role"] == "tsundere" {
			return "tsundere", nil
		}
		if in["role"] == "cute" {
			return "cute", nil
		}
		return "tsundere", nil
	}, map[string]bool{"tsundere": true, "cute": true}))

	// 链接节点
	err = g.AddEdge(compose.START, "lambda")
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("tsundere", "model")
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("cute", "model")
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("model", compose.END)
	if err != nil {
		panic(err)
	}
	return g

}
func OutSideOrcGraph(ctx context.Context) *compose.Graph[map[string]string, string] {
	insideGraph := Graph(ctx)
	//外部图
	outsideGraph := compose.NewGraph[map[string]string, string]()
	//创建节点
	lambda := compose.InvokableLambda(func(ctx context.Context, input map[string]string) (output map[string]string, err error) {
		return input, nil
	})
	writeLambda := compose.InvokableLambda(func(ctx context.Context, input *schema.Message) (output string, err error) {
		f, err := os.OpenFile("orc_graph_withgraph.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return "", err
		}
		defer f.Close()
		if _, err := f.WriteString(input.Content + "\n---\n"); err != nil {
			return "", err
		}
		return "已经写入文件，请前往文件内查看内容", nil
	})
	//添加节点
	err := outsideGraph.AddLambdaNode("lambda", lambda)
	if err != nil {
		panic(err)
	}
	err = outsideGraph.AddGraphNode("inside", insideGraph)
	if err != nil {
		panic(err)
	}
	err = outsideGraph.AddLambdaNode("write", writeLambda)
	if err != nil {
		panic(err)
	}
	//链接节点
	err = outsideGraph.AddEdge(compose.START, "lambda")
	if err != nil {
		panic(err)
	}
	err = outsideGraph.AddEdge("lambda", "inside")
	if err != nil {
		panic(err)
	}
	err = outsideGraph.AddEdge("inside", "write")
	if err != nil {
		panic(err)
	}
	err = outsideGraph.AddEdge("write", compose.END)
	if err != nil {
		panic(err)
	}
	return outsideGraph
}
