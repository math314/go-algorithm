package bst

import (
	"math/rand"
	"sort"
	"testing"
)

func TestAvlTree_Insert(t *testing.T) {
	testVals := make([]int, 0)
	for i := 0; i < 10000; i++ {
		testVals = append(testVals , rand.Int())
	}

	avlTree := &AvlTree{nil, 0}

	for i := 0; i < len(testVals); i++ {
		tmp := ComparableInt(testVals[i])
		avlTree.Insert(&tmp)

		if avlTree.size != i + 1 {
			t.Fatal("size mismatch")
		}

		sorted := testVals[:avlTree.size]
		sort.Ints(sorted)

		for idx, val := range avlTree.toList() {
			tmp, _ := val.(*ComparableInt)
			if int(*tmp) != sorted[idx] {
				t.Fatal("value is not sorted")
			}
		}
	}
}

func TestAvlTree_Delete(t *testing.T) {
	testVals := make([]int, 0)
	for i := 0; i < 10000; i++ {
		testVals = append(testVals , rand.Int())
	}

	avlTree := &AvlTree{nil, 0}

	for i := 0; i < len(testVals); i++ {
		tmp := ComparableInt(testVals[i])
		avlTree.Insert(&tmp)
	}

	rand.Shuffle(len(testVals), func(i, j int) {
		testVals[i], testVals[j] = testVals[j], testVals[i]
	})

	for i := 0; i < len(testVals); i++ {
		tmp := ComparableInt(testVals[i])
		avlTree.Delete(&tmp)

		if avlTree.size != len(testVals) - i - 1 {
			t.Fatal("size mismatch")
		}

		sorted := testVals[i+1:]
		sort.Ints(sorted)

		for idx, val := range avlTree.toList() {
			tmp, _ := val.(*ComparableInt)
			if int(*tmp) != sorted[idx] {
				t.Fatal("value is not sorted")
			}
		}
	}

}