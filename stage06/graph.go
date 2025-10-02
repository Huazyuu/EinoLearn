package stage06

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/compose"
)

func Graph(choice string) {
	ctx := context.Background()
	g := compose.NewGraph[string, string]()
	lambda0 := compose.InvokableLambda(func(ctx context.Context, input string) (output string, err error) {
		switch input {
		case "1":
			output = "毫猫"
		case "2":
			output = "耄耋"
		case "3":
			output = "哈基汪"
		default:
			output = "我不理解你的意思"
		}
		return output, nil
	})
	lambda1 := compose.InvokableLambda(func(ctx context.Context, input string) (output string, err error) {
		return "喵喵喵喵喵喵喵", nil
	})
	lambda2 := compose.InvokableLambda(func(ctx context.Context, input string) (output string, err error) {
		return "哈!", nil
	})
	lambda3 := compose.InvokableLambda(func(ctx context.Context, input string) (output string, err error) {
		return "大旋风!!!", nil
	})
	err := g.AddLambdaNode("lambda0", lambda0)
	if err != nil {
		panic(err)
	}
	err = g.AddLambdaNode("lambda1", lambda1)
	if err != nil {
		panic(err)
	}
	err = g.AddLambdaNode("lambda2", lambda2)
	if err != nil {
		panic(err)
	}
	err = g.AddLambdaNode("lambda3", lambda3)
	if err != nil {
		panic(err)
	}
	err = g.AddBranch("lambda0",
		compose.NewGraphBranch(
			func(ctx context.Context, in string) (endNode string, err error) {
				switch in {
				case "毫猫":
					return "lambda1", nil
				case "耄耋":
					return "lambda2", nil
				case "哈基汪":
					return "lambda3", nil
				}
				// 否则，返回 compose.END，表示流程结束
				return compose.END, nil
			}, map[string]bool{
				"lambda1": true,
				"lambda2": true,
				"lambda3": true,
			}))
	if err != nil {
		panic(err)
	}
	err = g.AddEdge(compose.START, "lambda0")
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("lambda1", compose.END)
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("lambda2", compose.END)
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("lambda3", compose.END)
	if err != nil {
		panic(err)
	}
	// 编译
	r, err := g.Compile(ctx)
	if err != nil {
		panic(err)
	}
	// 执行
	answer, err := r.Invoke(ctx, choice)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)

}
