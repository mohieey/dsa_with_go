package main

import (
	"dsa_with_go/dswithgo"
	"fmt"
)

func main() {
	values := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}

	heap := dswithgo.MaxHeap{}
	fmt.Println(heap)

	for _, v := range values {
		heap.Insert(v)
		fmt.Println(heap)
	}

	for i := 0; i < 5; i++ {

		heap.Extract()
		fmt.Println(heap)
	}

}
