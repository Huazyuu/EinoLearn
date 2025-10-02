package main

import (
	"einolearn/demo01/stage04_RAG/indexer"
	"einolearn/demo01/stage04_RAG/retriever"
	"einolearn/demo01/stage04_RAG/trans"
	"fmt"
)

func main() {
	// 分词分块
	docs := trans.TransDoc("./document.md")
	// 索引
	indexer.IndexerRAG(docs)
	// 检索
	res := retriever.RetrieveRAG("太原理工", 3)
	for _, doc := range res {
		fmt.Println(doc.ID)

		fmt.Println(doc)
	}
	fmt.Println("------------------")
	res = retriever.RetrieveRAG("元神主角是谁", 3)
	for _, doc := range res {
		fmt.Println(doc.ID)
		fmt.Println(doc)
	}
	fmt.Println("------------------")
	res = retriever.RetrieveRAG("tyut", 3)
	for _, doc := range res {
		fmt.Println(doc.ID)

		fmt.Println(doc)
	}
	fmt.Println("------------------")
	res = retriever.RetrieveRAG("软件工程", 3)
	for _, doc := range res {
		fmt.Println(doc.ID)
		fmt.Println(doc)
	}
	fmt.Println("------------------")
	res = retriever.RetrieveRAG("se", 3)
	for _, doc := range res {
		fmt.Println(doc.ID)
		fmt.Println(doc)
	}
}
