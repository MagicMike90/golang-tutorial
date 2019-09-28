package main

import (
	"fmt"
)

type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil, nil}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if t.Next == nil {
		temp := t
		t.Next = &Node{v, temp, nil}
		return -2
	}

	return addNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func reverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	// go to the end of the doubly linked list
	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}

	for temp.Previous != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}
	fmt.Printf("%d -> ", temp.Value)
	fmt.Println()
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookupNode(t *Node, v int) bool {
	if root == nil {
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

var root = new(Node)

/**
PROs:
	1.traverse them in any direction you want and also you can insert and delete elements from a doubly linked list more easily
	2.if you lose the pointer to the head of a doubly linked list, you can still find the head node of that list
CONs:
	1.maintaining two pointers for each node
example:
	music player might be using a doubly linked list to represent your current list of songs and be able to go to the previous song as well as the next one!
*/
func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, 1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 0)
	addNode(root, 0)
	traverse(root)
	addNode(root, 100)
	fmt.Println("Size:", size(root))
	traverse(root)
	reverse(root)
}
