package indexer

import (
	"context"
	"demo01/stage04_RAG/milvusCli"
	"log"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/indexer/milvus"
	"github.com/cloudwego/eino/schema"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

var Collection = "Demo01RAGCollection_20251001_TYUT_SE"

var fields = []*entity.Field{
	{
		Name:     "id",
		DataType: entity.FieldTypeVarChar,
		TypeParams: map[string]string{
			"max_length": "255",
		},
		PrimaryKey: true,
	},
	{
		Name:     "vector", // 确保字段名匹配
		DataType: entity.FieldTypeBinaryVector,
		TypeParams: map[string]string{
			"dim": "81920",
		},
	},
	{
		Name:     "content",
		DataType: entity.FieldTypeVarChar,
		TypeParams: map[string]string{
			"max_length": "8192",
		},
	},
	{
		Name:     "metadata",
		DataType: entity.FieldTypeJSON,
	},
}

func IndexerRAG(docs []*schema.Document) {
	ctx := context.Background()
	timeout := time.Second * 30
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		Timeout: &timeout,
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   "doubao-embedding-text-240715",
	})
	if err != nil {
		panic(err)
	}
	indexer, err := milvus.NewIndexer(ctx, &milvus.IndexerConfig{
		Client:     milvusCli.MilvusCli,
		Collection: Collection,
		Fields:     fields,
		Embedding:  embedder,
	})
	if err != nil {
		log.Fatalf("Failed to create indexer: %v", err)
	}
	for _, doc := range docs {
		storeDoc := []*schema.Document{
			{
				ID:       doc.ID,
				Content:  doc.Content,
				MetaData: doc.MetaData,
			},
		}
		ids, err := indexer.Store(ctx, storeDoc)
		if err != nil {
			log.Fatalf("Failed to store document: %v", err)
		}
		log.Printf("Stored document with ID: %v", ids)
	}
}
