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

func (t *Tree) BFS() []int {
	var currentNode *Node
	vals := []int{}
	queue := []*Node{t.Root}

	if t.Root == nil {
		return vals
	}

	for len(queue) > 0 {
		currentNode = queue[0]
		queue = queue[1:]
		vals = append(vals, currentNode.Value)
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}
	return vals
}

func (t *Tree) DFSInOrder() []int {
	values := []int{}
	if t.Root != nil {
		return t.traverseInOrder(t.Root, &values)
	} else {
		return values
	}

}

func (t *Tree) DFSPreOrder() []int {
	values := []int{}
	if t.Root != nil {
		return t.traversePreOrder(t.Root, &values)
	} else {
		return values
	}

}

func (t *Tree) DFSPostOrder() []int {
	values := []int{}
	if t.Root != nil {
		return t.traversePostOrder(t.Root, &values)
	} else {
		return values
	}

}

func (t *Tree) traverseInOrder(node *Node, values *[]int) []int {
	if node.Left != nil {
		t.traverseInOrder(node.Left, values)
	}
	*values = append(*values, node.Value)
	if node.Right != nil {
		t.traverseInOrder(node.Right, values)
	}

	return *values
}

func (t *Tree) traversePreOrder(node *Node, values *[]int) []int {
	*values = append(*values, node.Value)
	if node.Left != nil {
		t.traversePreOrder(node.Left, values)
	}
	if node.Right != nil {
		t.traversePreOrder(node.Right, values)
	}

	return *values
}

func (t *Tree) traversePostOrder(node *Node, values *[]int) []int {
	if node.Left != nil {
		t.traversePostOrder(node.Left, values)
	}
	if node.Right != nil {
		t.traversePostOrder(node.Right, values)
	}
	*values = append(*values, node.Value)

	return *values
}
