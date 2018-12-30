package main

import "fmt"

type Node struct {
	value       int64
	left_child  *Node
	right_child *Node
}

func NewNode() *Node {
	node := Node{}
	return &node
}

func (root *Node) insert(value int64) {
	if (Node{}) == *root {
		*root = *NewNode()
		root.value = value
		root.left_child = NewNode()
		root.right_child = NewNode()

	} else if value <= root.value {
		root.left_child.insert(value)

	} else {
		root.right_child.insert(value)
	}
}

func (root Node) preview(position string, level int) {
	// in-order preview
	if (Node{}) != root {
		fmt.Printf("position: %s\nlevel: %d\nvalue: %d\n", position, level, root.value)
		root.left_child.preview("left", level + 1)
		root.right_child.preview("right", level + 1)
	}
}

func main() {
	root := Node{}
	child := int64(123)
	child_2 := int64(1234)
	child_3 := int64(122)
	child_4 := int64(123)

	root.insert(child)
	root.insert(child_2)
	root.insert(child_3)
	root.insert(child_4)
	root.preview("root", 0)
}
