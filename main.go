package main

import (
	"dsa_with_go/dswithgo"
	"fmt"
)

func main() {

	stack := dswithgo.Stack{}
	stack.Print()

	values := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range values {
		stack.Push(v)
	}

	stack.Print()

	fmt.Println(stack.Pop())
	stack.Print()

}
