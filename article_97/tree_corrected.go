//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

import "unsafe"
import "github.com/Konstantin8105/c4go/noarch"

// Node - transpiled function from  /root/tree.c:5
type Node struct {
	left  []Node
	right []Node
	value []byte
}

// insert_new_node - transpiled function from  /root/tree.c:12
func insert_new_node(root [][]Node, value []byte) {
	var cmp int32
	if len(root[0]) == 0 {
		root[0] = make([]Node, 1)
		(root[0])[0].value = make([]byte, noarch.Strlen(value))
		noarch.Strcpy((root[0])[0].value, value)
		(root[0])[0].left = nil
		(root[0])[0].right = nil
		return
	}
	cmp = noarch.Strcmp(value, (root[0])[0].value)
	if cmp < 0 {
		insert_new_node((*[1000000][]Node)(unsafe.Pointer(&(root[0])[0].left))[:], value)
	} else {
		insert_new_node((*[1000000][]Node)(unsafe.Pointer(&(root[0])[0].right))[:], value)
	}
}

// traverse_tree - transpiled function from  /root/tree.c:36
func traverse_tree(root []Node, callback_function func([]byte)) {
	if len(root) == 0 {
		return
	}
	traverse_tree(root[0].left, callback_function)
	callback_function(root[0].value)
	traverse_tree(root[0].right, callback_function)
}

// callback_function - transpiled function from  /root/tree.c:47
func callback_function(value []byte) {
	noarch.Printf([]byte("%s\n\x00"), value)
}

// main - transpiled function from  /root/tree.c:52
func main() {
	defer noarch.AtexitRun()
	var root []Node
	insert_new_node((*[1000000][]Node)(unsafe.Pointer(&root))[:], []byte("xxx\x00"))
	insert_new_node((*[1000000][]Node)(unsafe.Pointer(&root))[:], []byte("aaa\x00"))
	insert_new_node((*[1000000][]Node)(unsafe.Pointer(&root))[:], []byte("bbb\x00"))
	insert_new_node((*[1000000][]Node)(unsafe.Pointer(&root))[:], []byte("ccc\x00"))
	insert_new_node((*[1000000][]Node)(unsafe.Pointer(&root))[:], []byte("yyy\x00"))
	insert_new_node((*[1000000][]Node)(unsafe.Pointer(&root))[:], []byte("yyy\x00"))
	traverse_tree(root, callback_function)
	return
}
