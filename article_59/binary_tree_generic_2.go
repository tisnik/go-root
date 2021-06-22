// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá osmá část
//    Generické datové typy v jazyce Go?
//    https://www.root.cz/clanky/genericke-datove-typy-v-jazyce-go/

package main

import (
	"fmt"
	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type Node struct {
	Value Item
	Left  *Node
	Right *Node
}

type ItemBinaryTree struct {
	Root *Node
}

func (bt *ItemBinaryTree) Insert(value Item) {
	node := &ItemNode{value, nil, nil}
	if bt.Root == nil {
		bt.Root = node
	} else {
		insertItemNode(bt.Root, node)
	}
}

func insertItemNode(node, newNode *Node) {
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertItemNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertItemNode(node.Right, newNode)
		}
	}
}

func printItemTree(node *Node, level int) {
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		printItemTree(node.Left, level)
		fmt.Printf(format+"%v\n", node.Value)
		printItemTree(node.Right, level)
	}
}

/*
func main() {
	var bt ItemBinaryTree
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
*/
