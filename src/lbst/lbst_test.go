package lbst

import (
	"math/rand"
	"sort"
	"testing"
)

func TestLogLevel(t *testing.T) {
	logVals := []int{0, 1, 2, 2, 3, 3, 3, 3, 4}
	for i, v := range logVals {
		if logLevel(i) != v {
			t.Fatal("loglevel mismatch")
		}
	}
}

func TestLogSmaller(t *testing.T) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			smaller1 := logLevel(i) < logLevel(j)
			smaller2 := logSmaller(i, j)
			if smaller1 != smaller2 {
				t.Fatal("logSmaller returned a wrong value")
			}
		}
	}
}

func verifyTree(t *testing.T, tree *LogarithmicBst, v *node) {
	if v == nil {
		return
	}

	ls, rs := size(v.l), size(v.r)
	if size(v) != ls+rs+1 {
		t.Fatal("invalid size")
	}

	lDiff := logLevel(ls) - logLevel(rs)
	if lDiff < -1 || lDiff > 1 {
		t.Fatal("invalid balance")
	}

	if v.l != nil {
		if v.l.p != v {
			t.Fatal("parent mismatch")
		}
		if tree.comp(v.l.val, v.val) != -1 {
			t.Fatal("invalid order")
		}
		verifyTree(t, tree, v.l)
	}
	if v.r != nil {
		if v.r.p != v {
			t.Fatal("parent mismatch")
		}
		if tree.comp(v.val, v.r.val) != -1 {
			t.Fatal("invalid order")
		}
		verifyTree(t, tree, v.r)
	}
}

func randomInts(len int) []int {
	mp := map[int]struct{}{}
	ret := make([]int, 0, len)
	for i := 0; i < len; i++ {
		var x int
		for {
			x = rand.Int()
			if _, found := mp[x]; !found {
				break
			}
		}
		ret = append(ret, x)
	}
	return ret
}

func TestLbst_Insert(t *testing.T) {
	testVals := randomInts(1000)
	tree := NewLogarithmicBst(IntComparator)

	for i := 0; i < len(testVals); i++ {
		tree.Insert(testVals[i])

		verifyTree(t, tree, tree.root)

		if tree.Size() != i+1 {
			t.Fatal("size mismatch")
		}

		sorted := testVals[:tree.Size()]
		sort.Ints(sorted)

		idx := 0
		for it := tree.Iterator(); it.Next(); {
			val := it.Value().(int)
			if sorted[idx] != val {
				t.Fatal("value is not sorted")
			}
			idx++
		}
	}
}

func TestLbst_Find(t *testing.T) {
	testVals := randomInts(1000 * 2)
	tree := NewLogarithmicBst(IntComparator)
	for i := 0; i < 1000; i++ {
		tree.Insert(testVals[i])
	}
	for i, val := range testVals {
		found := tree.Find(val)
		expected := i < 1000
		if found != expected {
			t.Fatal("found mismatch")
		}
	}
}

func TestLbst_LessCount(t *testing.T) {
	testVals := randomInts(1000)
	tree := NewLogarithmicBst(IntComparator)
	for i := 0; i < 1000; i++ {
		tree.Insert(testVals[i])
	}
	sort.Ints(testVals)
	for i, val := range testVals {
		i2 := tree.LowerCount(val)
		if i != i2 {
			t.Fatal("index mismatch")
		}
		i2PlusOne := tree.LowerCount(val + 1)
		if i2+1 != i2PlusOne {
			t.Fatal("index mismatch")
		}
	}
}

func TestLbst_Delete(t *testing.T) {
	testVals := randomInts(1000)
	tree := NewLogarithmicBst(IntComparator)
	for _, val := range testVals {
		tree.Insert(val)
	}
	rand.Shuffle(len(testVals), func(i, j int) {
		testVals[i], testVals[j] = testVals[j], testVals[i]
	})

	for i := 0; i < len(testVals); i++ {
		tree.Delete(testVals[i])

		verifyTree(t, tree, tree.root)

		if tree.Size() != len(testVals)-1-i {
			t.Fatal("size mismatch")
		}

		sorted := append(testVals[:0:0], testVals[i+1:]...)
		sort.Ints(sorted)

		idx := 0
		for it := tree.Iterator(); it.Next(); {
			val := it.Value().(int)
			if sorted[idx] != val {
				t.Fatal("value is not sorted")
			}
			idx++
		}
	}
}

func TestLbst_DuplicatedInsert(t *testing.T) {
	testVals := randomInts(1000)
	tree := NewLogarithmicBst(IntComparator)
	for i, val := range testVals {
		if inserted := tree.Insert(val); !inserted {
			t.Fatal("not inserted but returned True")
		}
		if tree.Size() != i+1 {
			t.Fatal("size mismatch")
		}
		verifyTree(t, tree, tree.root)
	}
	for _, val := range testVals {
		if inserted := tree.Insert(val); inserted {
			t.Fatal("already inserted but returned False")
		}
		if tree.Size() != len(testVals) {
			t.Fatal("size mismatch")
		}
		verifyTree(t, tree, tree.root)
	}

}
