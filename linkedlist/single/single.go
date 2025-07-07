package single

import "fmt"

type Node struct {
	val int
	next *Node
}

func NewList() *Node {
	return &Node{0, nil}
}

func Add(n *Node, val int) {

	tmp := n

	for tmp.next != nil {
		tmp = tmp.next
	}
	
	tmp.next = &Node{val, nil}
}

func Del(n *Node, val int) {
	tmp := n
	
	var target *Node

	for tmp != nil {
		if tmp.next != nil && tmp.next.val == val {
			target = tmp	
			break
		}
		tmp = tmp.next
	}

	if target != nil {
		next := target.next.next
		target.next = next
	}
}

func Traverse(node *Node) {
	tmp := node.next
	for tmp != nil {
		fmt.Println(tmp.val)
		tmp = tmp.next
	}
}
