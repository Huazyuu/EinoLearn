package trans

import (
	"context"
	"os"
	"strconv"

	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
)

func TransDoc(path string) []*schema.Document {
	ctx := context.Background()
	splitter, err := markdown.NewHeaderSplitter(ctx, &markdown.HeaderConfig{
		Headers: map[string]string{
			"#":   "h1",
			"##":  "h2",
			"###": "h3",
		},
		TrimHeaders: false,
	})
	if err != nil {
		panic(err)
	}
	content, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer content.Close()
	bs, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	docs := []*schema.Document{
		{
			ID:      uuid.New().String(),
			Content: string(bs),
		},
	}
	results, err := splitter.Transform(ctx, docs)
	if err != nil {
		panic(err)
	}
	for i, doc := range results {
		doc.ID = docs[0].ID + "_" + strconv.Itoa(i)
		println(doc.ID)
	}
	return results

}
