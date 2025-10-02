package main

import (
	"context"
	"einolearn/demo01/stage07"
	"fmt"
)

func main() {
	ctx := context.Background()
	g := stage07.OutSideOrcGraph(ctx)
	r, err := g.Compile(ctx)
	if err != nil {
		panic(err)
	}
	result, err := r.Invoke(ctx, map[string]string{"role": "tsundere", "content": "嘿嘿嘿,你好啊小妹妹要不要和叔叔一起玩元神"})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
