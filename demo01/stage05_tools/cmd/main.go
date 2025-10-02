package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/eino-ext/components/tool/browseruse"
)

func main() {
	browserTool, err := browseruse.NewBrowserUseTool(context.Background(), &browseruse.Config{})
	if err != nil {
		log.Fatal(err)
	}
	url := `https://ys.mihoyo.com/tool`
	res, err := browserTool.Execute(&browseruse.Param{
		Action: browseruse.ActionGoToURL,
		URL:    &url,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	time.Sleep(10 * time.Second)
	browserTool.Cleanup()
}
