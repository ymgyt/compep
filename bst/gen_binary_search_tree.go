// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package bst

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

// cat binary_search_tree.go | genny gen ValueType=string

// IntZeroValue -
var IntZeroValue int

// NewIntBinarySearchTree -
func NewIntBinarySearchTree() *IntBinarySearchTree {
	return &IntBinarySearchTree{}
}

// IntNode -
type IntNode struct {
	key   int
	value int

	right *IntNode
	left  *IntNode
}

// IntBinarySearchTree -
type IntBinarySearchTree struct {
	root *IntNode
	sync.RWMutex
}

// Insert insert given key-value. if given value v is already exists in tree, just discard it.
func (bst *IntBinarySearchTree) Insert(k int, v int) {
	bst.Lock()
	defer bst.Unlock()

	n := &IntNode{k, v, nil, nil}

	if bst.root == nil {
		bst.root = n
	}
	bst.insertIntNode(bst.root, n)
}

func (bst *IntBinarySearchTree) insertIntNode(node, newIntNode *IntNode) {
	if newIntNode.key < node.key {
		if node.left == nil {
			node.left = newIntNode
		} else {
			bst.insertIntNode(node.left, newIntNode)
		}
	} else if newIntNode.key > node.key {
		if node.right == nil {
			node.right = newIntNode
		} else {
			bst.insertIntNode(node.right, newIntNode)
		}
	}
	// same value discard it
}

// Search search given key. if not found, return false.
func (bst *IntBinarySearchTree) Search(key int) (int, bool) {
	bst.RLock()
	defer bst.RUnlock()
	return bst.search(bst.root, key)
}

func (bst *IntBinarySearchTree) search(n *IntNode, key int) (int, bool) {
	if n == nil {
		return IntZeroValue, false
	}
	if key < n.key {
		return bst.search(n.left, key)
	} else if key > n.key {
		return bst.search(n.right, key)
	}
	return n.value, true
}

func (bst *IntBinarySearchTree) String() string {
	bst.Lock()
	defer bst.Unlock()
	var b bytes.Buffer
	bst.stringify(&b, bst.root, 0)
	return b.String()
}

func (bst *IntBinarySearchTree) stringify(w io.Writer, n *IntNode, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "  "
		}
		format += "---["
		level++
		bst.stringify(w, n.left, level)
		fmt.Fprintf(w, format+"%v\n", n.key)
		bst.stringify(w, n.right, level)
	}
}

// cat binary_search_tree.go | genny gen ValueType=string

// StringZeroValue -
var StringZeroValue string

// NewStringBinarySearchTree -
func NewStringBinarySearchTree() *StringBinarySearchTree {
	return &StringBinarySearchTree{}
}

// StringNode -
type StringNode struct {
	key   int
	value string

	right *StringNode
	left  *StringNode
}

// StringBinarySearchTree -
type StringBinarySearchTree struct {
	root *StringNode
	sync.RWMutex
}

// Insert insert given key-value. if given value v is already exists in tree, just discard it.
func (bst *StringBinarySearchTree) Insert(k int, v string) {
	bst.Lock()
	defer bst.Unlock()

	n := &StringNode{k, v, nil, nil}

	if bst.root == nil {
		bst.root = n
	}
	bst.insertStringNode(bst.root, n)
}

func (bst *StringBinarySearchTree) insertStringNode(node, newStringNode *StringNode) {
	if newStringNode.key < node.key {
		if node.left == nil {
			node.left = newStringNode
		} else {
			bst.insertStringNode(node.left, newStringNode)
		}
	} else if newStringNode.key > node.key {
		if node.right == nil {
			node.right = newStringNode
		} else {
			bst.insertStringNode(node.right, newStringNode)
		}
	}
	// same value discard it
}

// Search search given key. if not found, return false.
func (bst *StringBinarySearchTree) Search(key int) (string, bool) {
	bst.RLock()
	defer bst.RUnlock()
	return bst.search(bst.root, key)
}

func (bst *StringBinarySearchTree) search(n *StringNode, key int) (string, bool) {
	if n == nil {
		return StringZeroValue, false
	}
	if key < n.key {
		return bst.search(n.left, key)
	} else if key > n.key {
		return bst.search(n.right, key)
	}
	return n.value, true
}

func (bst *StringBinarySearchTree) String() string {
	bst.Lock()
	defer bst.Unlock()
	var b bytes.Buffer
	bst.stringify(&b, bst.root, 0)
	return b.String()
}

func (bst *StringBinarySearchTree) stringify(w io.Writer, n *StringNode, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "  "
		}
		format += "---["
		level++
		bst.stringify(w, n.left, level)
		fmt.Fprintf(w, format+"%v\n", n.key)
		bst.stringify(w, n.right, level)
	}
}
