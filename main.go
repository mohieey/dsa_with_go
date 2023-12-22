package main

import (
	"dsa_with_go/tree"
	"fmt"
)

func main() {
	tree := tree.Tree{}

	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(20)
	tree.Insert(170)
	tree.Insert(15)
	tree.Insert(1)

	fmt.Println(tree.Lookup(4))
}
