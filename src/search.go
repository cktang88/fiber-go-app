package main

import (
	"fmt"
	"log"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
	_ "github.com/mattn/go-sqlite3"
)

// global redis client
var redisClient *redisearch.Client

func redisInit() {
	// Create a client. By default a client is schemaless
	// unless a schema is provided when creating the index
	redisClient = redisearch.NewClient("localhost:6379", "myIndex")

	// Create a schema
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("body")).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewNumericField("date"))

	// Drop an existing index. If the index does not exist an error is returned
	// c.Drop()

	// Create the index with the given schema
	if err := redisClient.CreateIndex(sc); err != nil {
		log.Fatal(err)
	}
	// Create a document with an id and given score
	addText(redisClient, "doc1", "Hello world", "foo bar", 1.0)
	addText(redisClient, "doc2", "goodbye world", "asdf random body", 0.8)
}

func addText(c *redisearch.Client, id string, title string, body string, score float32) {
	// add doc
	doc := redisearch.NewDocument(id, score)
	doc.Set("title", title).
		Set("body", body).
		Set("date", time.Now().Unix())
	// Index the document. The API accepts multiple documents at a time
	if err := c.Index([]redisearch.Document{doc}...); err != nil {
		log.Fatal(err)
	}
}

func searchText(c *redisearch.Client, text string) {
	// Searching with limit and sorting
	docs, total, err := c.Search(redisearch.NewQuery(text).
		Limit(0, 2).
		SetReturnFields("title", "body"))

	fmt.Println(docs[0].Id, docs[0].Properties["title"], total, err)
	// Output: doc1 Hello world 1 <nil>
}
