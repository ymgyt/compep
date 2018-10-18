package bst

import (
	"bytes"
	"fmt"
	"io"
	"sync"

	"github.com/cheekybits/genny/generic"
)

//go:generate genny -in $GOFILE -out gen_$GOFILE gen Item=int,string

// cat binary_search_tree.go | genny gen ValueType=string

// Item -
type Item generic.Type

// ItemZeroValue -
var ItemZeroValue Item

// NewItemBinarySearchTree -
func NewItemBinarySearchTree() *ItemBinarySearchTree {
	return &ItemBinarySearchTree{}
}

// ItemNode -
type ItemNode struct {
	key   int
	value Item

	right *ItemNode
	left  *ItemNode
}

// ItemBinarySearchTree -
type ItemBinarySearchTree struct {
	root *ItemNode
	sync.RWMutex
}

// Insert insert given key-value. if given value v is already exists in tree, just discard it.
func (bst *ItemBinarySearchTree) Insert(k int, v Item) {
	bst.Lock()
	defer bst.Unlock()

	n := &ItemNode{k, v, nil, nil}

	if bst.root == nil {
		bst.root = n
	}
	bst.insertItemNode(bst.root, n)
}

func (bst *ItemBinarySearchTree) insertItemNode(node, newItemNode *ItemNode) {
	if newItemNode.key < node.key {
		if node.left == nil {
			node.left = newItemNode
		} else {
			bst.insertItemNode(node.left, newItemNode)
		}
	} else if newItemNode.key > node.key {
		if node.right == nil {
			node.right = newItemNode
		} else {
			bst.insertItemNode(node.right, newItemNode)
		}
	}
	// same value discard it
}

// Search search given key. if not found, return false.
func (bst *ItemBinarySearchTree) Search(key int) (Item, bool) {
	bst.RLock()
	defer bst.RUnlock()
	return bst.search(bst.root, key)
}

func (bst *ItemBinarySearchTree) search(n *ItemNode, key int) (Item, bool) {
	if n == nil {
		return ItemZeroValue, false
	}
	if key < n.key {
		return bst.search(n.left, key)
	} else if key > n.key {
		return bst.search(n.right, key)
	}
	return n.value, true
}

func (bst *ItemBinarySearchTree) String() string {
	bst.Lock()
	defer bst.Unlock()
	var b bytes.Buffer
	bst.stringify(&b, bst.root, 0)
	return b.String()
}

func (bst *ItemBinarySearchTree) stringify(w io.Writer, n *ItemNode, level int) {
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
