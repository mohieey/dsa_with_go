package main

import (
	"dsa_with_go/dswithgo"
	"fmt"
)

func main() {

	linkedList := dswithgo.LinkedList{}

	values := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range values {
		linkedList.Prepend(v)
	}

	linkedList.Print()
	fmt.Println(linkedList.Length)

	fmt.Println(linkedList.RemoveValue(9))
	fmt.Println(linkedList.RemoveValue(55))
	fmt.Println(linkedList.RemoveValue(17))

	linkedList.Print()
	fmt.Println(linkedList.Length)

}
