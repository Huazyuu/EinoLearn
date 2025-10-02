package main

import (
	"context"
	"einolearn/demo01/stage07"
	"fmt"
)

func main() {
	ctx := context.Background()
	// 测试傲娇角色
	tsundereInput := map[string]string{
		"role":    "tsundere",
		"content": "原神启动!",
	}
	stage07.GraphWithState(ctx, tsundereInput)
	fmt.Println("=================================")

	// 测试可爱角色
	cuteInput := map[string]string{
		"role":    "cute",
		"content": "原神启动!",
	}
	stage07.GraphWithState(ctx, cuteInput)
	fmt.Println("=================================")

	// 测试默认角色
	defaultInput := map[string]string{
		"role":    "user",
		"content": "你说得对,为什么不玩原神!",
	}
	stage07.GraphWithState(ctx, defaultInput)
}
