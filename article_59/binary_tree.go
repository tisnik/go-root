package main

import (
	"fmt"
)

type Item int

type Node struct {
	Value Item
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (bt *BinaryTree) Insert(value Item) {
	node := &Node{value, nil, nil}
	if bt.Root == nil {
		bt.Root = node
	} else {
		insertNode(bt.Root, node)
	}
}

func insertNode(node, newNode *Node) {
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNode(node.Right, newNode)
		}
	}
}

func printTree(node *Node, level int) {
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		printTree(node.Left, level)
		fmt.Printf(format+"%v\n", node.Value)
		printTree(node.Right, level)
	}
}

func main() {
	var bt BinaryTree
	bt.Insert(8)

	bt.Insert(3)
	bt.Insert(11)

	bt.Insert(1)
	bt.Insert(0)
	bt.Insert(2)

	bt.Insert(5)
	bt.Insert(4)
	bt.Insert(6)

	bt.Insert(9)
	bt.Insert(8)
	bt.Insert(10)

	bt.Insert(13)
	bt.Insert(12)
	bt.Insert(14)

	printTree(bt.Root, 0)
}
