package dswithgo

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Prepend(num int) {
	newNode := Node{num, l.Head}
	newNode.next = l.Head
	l.Head = &newNode
	l.Length++
}

func (l *LinkedList) RemoveValue(num int) bool {
	if l.Length == 0 {
		return false
	}

	if l.Head.value == num {

		l.Head = l.Head.next
		l.Length--
		return true
	}

	_, prevNode := l.Scan(num)
	if prevNode == nil {
		return false
	}

	prevNode.next = prevNode.next.next
	l.Length--
	return true
}

func (l *LinkedList) Scan(num int) (*Node, *Node) {
	var prevNode *Node
	currentNode := l.Head

	for currentNode != nil {
		if currentNode.value == num {
			return currentNode, prevNode
		}
		prevNode = currentNode
		currentNode = currentNode.next
	}

	return nil, nil
}

func (l *LinkedList) Print() {
	if l.Length == 0 {
		fmt.Println("the list is empty")
		return
	}

	currentNode := l.Head
	for currentNode != nil {
		if currentNode != l.Head {
			fmt.Print(" => ")
		}
		fmt.Print(currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Println()
}
