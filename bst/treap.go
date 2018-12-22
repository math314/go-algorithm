package bst

import (
	"fmt"
	"math/rand"
	"strings"
)

type treapNode struct {
	data ValType
	parent, left, right *treapNode
	treeSize int32
	prob int32
}

type Treap struct {
	root *treapNode
	size int
}

func recalcTreeSize(node *treapNode) {
	node.treeSize = 1
	if node.left != nil {
		node.treeSize += node.left.treeSize
	}
	if node.right != nil {
		node.treeSize += node.right.treeSize
	}
}

func treapRotateRight(node *treapNode) *treapNode {
	newRoot := node.left
	oldRoot := node

	if newRoot.right != nil {
		newRoot.right.parent = oldRoot
	}
	oldRoot.left = newRoot.right
	newRoot.right = oldRoot
	newRoot.parent = oldRoot.parent
	oldRoot.parent = newRoot

	recalcTreeSize(oldRoot)
	recalcTreeSize(newRoot)

	return newRoot
}

func treapRotateLeft(node *treapNode) *treapNode {
	newRoot := node.right
	oldRoot := node

	if newRoot.left != nil {
		newRoot.left.parent = oldRoot
	}
	oldRoot.right = newRoot.left
	newRoot.left = oldRoot
	newRoot.parent = oldRoot.parent
	oldRoot.parent = newRoot

	recalcTreeSize(oldRoot)
	recalcTreeSize(newRoot)

	return newRoot
}



func treapNodeInsert(node *treapNode, val ValType) (*treapNode, bool) {
	if node == nil {
		return &treapNode{val, nil, nil, nil, 1, rand.Int31()}, true
	}

	switch val.Compare(node.data) {
	case EQUAL:
		//duplication is not allowed
		return node, false
	case LESS:
		var inserted bool
		node.left, inserted = treapNodeInsert(node.left, val)
		node.left.parent = node
		if !inserted {
			return node, false
		}
		if node.prob > node.left.prob {
			node = treapRotateRight(node)
		} else {
			recalcTreeSize(node)
		}
	case GREATER:
		var inserted bool
		node.right, inserted = treapNodeInsert(node.right, val)
		node.right.parent = node
		if !inserted {
			return node, false
		}
		if node.prob > node.right.prob {
			node = treapRotateLeft(node)
		} else {
			recalcTreeSize(node)
		}
	}

	return node, true
}

func ValidateTreap(parent *treapNode, node *treapNode) int32 {
	if node == nil {
		return 0
	}
	childSize := int32(1)

	childSize += ValidateTreap(node, node.left)
	if parent != node.parent {
		panic("parent is incorrect")
	}
	if node.parent != nil {
		if node.parent.prob > node.prob {
			panic("prob ordering is incorrect")
		}
	}
	childSize += ValidateTreap(node, node.right)

	if childSize != node.treeSize {
		panic("treeSize is incorrect")
	}

	return childSize
}

func treapDebugPrint(node *treapNode, dep int) {
	if node == nil {
		return
	}
	treapDebugPrint(node.left, dep+ 1)
	a, _ := node.data.(*ComparableInt)
	fmt.Printf("%s%#v (prob = %d)\n", strings.Repeat(" ", dep), *a, node.prob)
	treapDebugPrint(node.right, dep + 1)
}

func treapWalkForToList(node *treapNode, ret *[]ValType) {
	if node == nil {
		return
	}
	treapWalkForToList(node.left, ret)
	*ret = append(*ret, node.data)
	treapWalkForToList(node.right, ret)
}

func NewTreap() *Treap {
	return &Treap{nil, 0}
}

func (treap *Treap) ToList() []ValType {
	ret := make([]ValType, 0, treap.size)
	treapWalkForToList(treap.root, &ret)
	return ret
}

func (treap *Treap) Insert(val ValType) {
	var inserted bool
	treap.root, inserted = treapNodeInsert(treap.root, val)
	if inserted {
		treap.size += 1
	}
}

func (treap *Treap) DebugPrint() {
	treapDebugPrint(treap.root, 0)
	ValidateTreap(nil, treap.root)
}
