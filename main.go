package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/crankshaft-agent/client"
	"github.com/crankshaft-agent/client/post"
	"log"
)

func main() {
	flag.Parse()

	c := client.Default

	params := &post.ListParams{Context: context.Background()}

	posts, err := c.Post.List(params)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range posts.Payload {
		fmt.Printf("PostID: %v UserID: %v Title: %v Body: %v\n", p.ID, p.UserID, p.Title, p.Body)
		fmt.Println("-------------------------------------------------------------------------------")
	}

	idParams := &post.GetParams{ID: 1, Context: context.Background()}
	post, err := c.Post.Get(idParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	fmt.Printf("PostID=%v UserID=%v Title=%v Body=%v\n", post.Payload.ID, post.Payload.UserID, post.Payload.Title, post.Payload.Body)

}
