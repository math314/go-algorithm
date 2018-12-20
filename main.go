package main

import (
	"fmt"
	"strings"
)

type CompareResult int
const (
	LESS CompareResult = iota - 1
	EQUAL
	GREATER
)

type Comparable interface {
	Compare(r interface{}) CompareResult
}

func ToCompareResult(diff int) CompareResult {
	switch {
	case diff < 0:
		return LESS
	case diff > 0:
		return GREATER
	default:
		return EQUAL
	}
}

type ValType Comparable

type avlNode struct {
	data ValType
	left *avlNode
	right *avlNode
	height int
}

type AvlTree struct {
	root *avlNode
	size int
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

func getHeight(node*avlNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func getBalance(node *avlNode) int {
	ret := 0
	if node.left != nil {
		ret += node.left.height
	}
	if node.right != nil {
		ret -= node.right.height
	}
	return ret
}

func rotateRight(node*avlNode) *avlNode {
	newRoot := node.left
	oldRoot := node

	oldRoot.left = newRoot.right
	newRoot.right = oldRoot

	oldRoot.height = 1 + max(getHeight(oldRoot.left), getHeight(oldRoot.right))
	newRoot.height = 1 + max(getHeight(newRoot.left), getHeight(newRoot.right))

	return newRoot
}

func rotateLeft(node*avlNode) *avlNode {
	newRoot := node.right
	oldRoot := node

	oldRoot.right = newRoot.left
	newRoot.left = oldRoot

	oldRoot.height = 1 + max(getHeight(oldRoot.left), getHeight(oldRoot.right))
	newRoot.height = 1 + max(getHeight(newRoot.left), getHeight(newRoot.right))

	return newRoot
}

func insert(node *avlNode, val ValType) (*avlNode, bool) {
	if node == nil {
		return &avlNode{val, nil, nil, 1}, true
	}

	switch val.Compare(node.data) {
	case EQUAL:
		//duplication is not allowed
		return node, false
	case LESS:
		var inserted bool
		node.left, inserted = insert(node.left, val)
		if !inserted {
			return node, false
		}
	case GREATER:
		var inserted bool
		node.right, inserted = insert(node.right, val)
		if !inserted {
			return node, false
		}
	}
	node.height = 1 + max(getHeight(node.left), getHeight(node.right))

	balance := getBalance(node)

	if balance > 1 {
		switch val.Compare(node.left.data) {
		case EQUAL:
		case LESS:
			// LL
			return rotateRight(node), true
		case GREATER:
			// LR
			node.left = rotateLeft(node.left)
			return rotateRight(node), true
		}
	} else if balance < -1 {
		switch val.Compare(node.right.data) {
		case EQUAL:
		case LESS:
			// RL
			node.right = rotateRight(node.right)
			return rotateLeft(node), true
		case GREATER:
			// RR
			return rotateLeft(node), true
		}
	} else {
		// balanced
		return node, true
	}

	panic("unreachable")
}

func debugPrint(node *avlNode) {
	if node == nil {
		return
	}
	debugPrint(node.left)
	 a, _ := node.data.(*ComparableInt)
	fmt.Printf("%s%#v\n", strings.Repeat(" ", node.height), *a)
	debugPrint(node.right)

}

func walkForToList(node *avlNode, ret []ValType) {
	if node == nil {
		return
	}
	walkForToList(node.left, ret)
	// slice size should be enough
	_ = append(ret, node.data)
	walkForToList(node.right, ret)
}

func (avlTree *AvlTree) toList() []ValType {
	ret := make([]ValType, 0, avlTree.size)
	walkForToList(avlTree.root, ret)
	return ret
}

func (avlTree *AvlTree) Insert(val ValType) {
	var inserted bool
	avlTree.root, inserted = insert(avlTree.root, val)
	if inserted {
		avlTree.size += 1
	}
}

type ComparableInt int

func (p *ComparableInt) Compare(r interface{}) CompareResult {
	rval, _ := r.(int)
	return ToCompareResult(int(*p) - rval)
}

func main() {
	avlTree := &AvlTree{nil, 0}

	for i := 0; i < 10; i++ {
		tmp := ComparableInt(i)
		avlTree.Insert(&tmp)
		debugPrint(avlTree.root)
	}

}
