package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/net/html"
	"os"
)

func main() {
	rootNode, err := html.Parse(os.Stdin)
	if err != nil {
		log.Error("parse error:", err)
		os.Exit(0)
	}
	fmt.Println(rootNode)
	fmt.Println(rootNode.Type)
	fmt.Println("---------------------")
	outline(nil, rootNode)
}

func outline(stack []string, node *html.Node) {

	if node.Type == html.ElementNode || node.Type == html.DocumentNode {
		stack = append(stack, node.Data)
		//fmt.Printf("%v", stack)
		fmt.Println(stack)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}

}
