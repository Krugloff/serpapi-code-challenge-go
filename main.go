package main

import (
	"fmt"
	"os"

	"serpapi-code-challenge-go/scanner_based/knowledge_graph"
)

func main() {
	data := html()

	// for i := 0; i < 1000; i++ {
		g := knowledge_graph.New(data)

		fmt.Printf("%q\n", g.JSON())
	// }
}

func html() []byte {
	data, err := os.ReadFile("files/van-gogh-paintings.html")

	if err != nil {
		fmt.Println("Error!", err)
	}

	return data
}