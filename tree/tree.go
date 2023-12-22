package tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(value int) {
	newNode := &Node{Value: value}

	if t.Root == nil {
		t.Root = newNode
		return
	}

	parent, currentNode := t.Root, t.Root
	for currentNode != nil {
		parent = currentNode
		if currentNode.Value < newNode.Value {
			currentNode = currentNode.Right
		} else {
			currentNode = currentNode.Left
		}
	}

	switch {
	case parent.Value < newNode.Value:
		parent.Right = newNode
	case parent.Value > newNode.Value:
		parent.Left = newNode

	}
}

func (t *Tree) Lookup(value int) *Node {
	if t.Root == nil {
		return nil
	}

	currentNode := t.Root
	for currentNode != nil {
		if currentNode.Value == value {
			return currentNode
		}

		if currentNode.Value < value {
			currentNode = currentNode.Right
		} else {
			currentNode = currentNode.Left
		}
	}
	return nil
}
