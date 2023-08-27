package main

import (
	"dsa_with_go/dswithgo"
	"fmt"
)

func main() {

	queue := dswithgo.Queue{}
	queue.Print()

	values := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range values {
		queue.Enqueue(v)
	}

	queue.Print()

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	queue.Print()

}
