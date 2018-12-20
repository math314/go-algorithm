package main


import (
	"math/rand"
	"sort"
	"testing"
)

func TestExampleSuccess(t *testing.T) {
	testVals := make([]int, 0)
	for i := 0; i < 100; i++ {
		testVals = append(testVals , rand.Int())
	}

	avlTree := &AvlTree{nil, 0}

	for i := 0; i < len(testVals); i++ {
		tmp := ComparableInt(testVals[i])
		avlTree.Insert(&tmp)

		if avlTree.size != i + 1 {
			t.Fatal("size mismatch")
		}

		sorted := make([]int, avlTree.size)
		sort.Ints(sorted)

		for idx, val := range avlTree.toList() {
			tmp, _ := val.(*ComparableInt)
			if int(*tmp) != sorted[idx] {
				t.Fatal("value is incorrect")
			}
		}
	}
}