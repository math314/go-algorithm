package main

import (
	"bst"
	"fmt"
)

func avlTest() {
	avlTree := bst.NewAvlTree()

	for i := 0; i < 10; i++ {
		tmp := bst.ComparableInt(i)
		avlTree.Insert(&tmp)
		avlTree.DebugPrint()
	}

	for i := 0; i < 10; i++ {
		tmp := bst.ComparableInt((i + 5) % 10)
		avlTree.Delete(&tmp)
		avlTree.DebugPrint()
	}
}

func treapTest() {
	treap := bst.NewTreap()

	for i := 0; i < 100; i++ {
		tmp := bst.ComparableInt(i)
		treap.Insert(&tmp)
		fmt.Printf("--- inserted %d ---\n", i)
		treap.DebugPrint()
	}

	//for i := 0; i < 10; i++ {
	//	tmp := bst.ComparableInt((i + 5) % 10)
	//	treap.Delete(&tmp)
	//	treap.DebugPrint()
	//}
}

func main() {

	// avlTest()
	treapTest()

}
