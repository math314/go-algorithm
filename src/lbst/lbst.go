package lbst

import (
	"math/bits"
	"strconv"
)

type node struct {
	l, r, p *node
	size    int
	val     interface{}
}

type LogarithmicBst struct {
	comp Comparator
	root *node
}

func NewLogarithmicBst(comp Comparator) *LogarithmicBst {
	return &LogarithmicBst{
		comp, nil,
	}
}

func logLevel(s int) int {
	return strconv.IntSize - bits.LeadingZeros(uint(s))
}

// Return true if logLevel(ls) < logLevel(rs) with O(1)
func logSmaller(ls, rs int) bool {
	if ls >= rs {
		return false
	}
	mixed := ls & rs
	return (mixed << 1) < rs
}

func size(t *node) int {
	if t == nil {
		return 0
	}
	return t.size
}

func rotLeft(t *node) *node {
	b := t.r
	t.r = b.l
	b.l = t
	b.p = t.p
	t.p = b
	if t.r != nil {
		t.r.p = t
	}
	b.size = t.size
	t.size = 1 + size(t.l) + size(t.r)
	return b
}

func rotRight(t *node) *node {
	a := t.l
	t.l = a.r
	a.r = t
	a.p = t.p
	t.p = a
	if t.l != nil {
		t.l.p = t
	}
	a.size = t.size
	t.size = 1 + size(t.l) + size(t.r)
	return a
}

func balance(t *node) *node {
	ls, rs := size(t.l), size(t.r)
	if ls < rs {
		// logLevel(ls) < logLevel(rs) - 1
		if logSmaller(ls, rs>>1) {
			r := t.r
			if logSmaller(size(r.r), size(r.l)) {
				t.r = rotRight(r)
			}
			t = rotLeft(t)
		}
	} else {
		// logLevel(ls) - 1 > logLevel(rs)
		if logSmaller(rs, ls>>1) {
			l := t.l
			if logSmaller(size(l.l), size(l.r)) {
				t.l = rotLeft(l)
			}
			t = rotRight(t)
		}
	}
	return t
}

func findMin(t *node) *node {
	for t.l != nil {
		t = t.l
	}
	return t
}

func (l *LogarithmicBst) insert(t *node, val interface{}) (*node, bool) {
	if t == nil {
		return &node{nil, nil, nil, 1, val}, true
	}
	var inserted bool
	comped := l.comp(val, t.val)
	if comped == -1 { // val < t.val
		t.l, inserted = l.insert(t.l, val)
		t.l.p = t
	} else if comped == 1 {
		t.r, inserted = l.insert(t.r, val)
		t.r.p = t
	} else {
		inserted = false
	}
	if inserted {
		t.size++
	}
	return balance(t), inserted
}

func (l *LogarithmicBst) Insert(val interface{}) bool {
	var inserted bool
	l.root, inserted = l.insert(l.root, val)
	return inserted
}

func (l *LogarithmicBst) find(val interface{}) *node {
	t := l.root
	for t != nil {
		comped := l.comp(val, t.val)
		if comped == 0 {
			return t
		}
		if comped == -1 {
			t = t.l
		} else if comped == 1 {
			t = t.r
		}
	}
	return nil
}

func (l *LogarithmicBst) Find(val interface{}) bool {
	foundNode := l.find(val)
	return foundNode != nil
}

// Return # of elements less than val
// O(log n)
func (l *LogarithmicBst) LowerCount(val interface{}) int {
	t := l.root
	ret := 0
	for t != nil {
		comped := l.comp(val, t.val)
		if comped == 0 {
			return ret + size(t.l)
		} else if comped == -1 {
			t = t.l
		} else if comped == 1 {
			ret += size(t.l) + 1
			t = t.r
		}
	}
	return ret
}

func (l *LogarithmicBst) delete(t *node, val interface{}) (*node, bool) {
	if t == nil {
		return nil, false
	}
	var deleted bool
	comped := l.comp(val, t.val)
	if comped == 0 {
		if t.r == nil {
			return t.l, true
		}
		minT := findMin(t.r)
		t.val, minT.val = minT.val, t.val
		t.r, deleted = l.delete(t.r, val)
		if t.r != nil {
			t.r.p = t
		}
	} else if comped == -1 { // val < t.val
		t.l, deleted = l.delete(t.l, val)
		if t.l != nil {
			t.l.p = t
		}
	} else if comped == 1 {
		t.r, deleted = l.delete(t.r, val)
		if t.r != nil {
			t.r.p = t
		}
	}
	if deleted {
		t.size--
	}

	return balance(t), deleted
}

func (l *LogarithmicBst) Delete(val interface{}) bool {
	var deleted bool
	l.root, deleted = l.delete(l.root, val)
	return deleted
}

func (l *LogarithmicBst) Size() int {
	return size(l.root)
}

type Iterator struct {
	l        *LogarithmicBst
	node     *node
	finished bool
}

func (l *LogarithmicBst) Iterator() *Iterator {
	return &Iterator{l, nil, false}
}

func (it *Iterator) Next() bool {
	if it.finished {
		return false
	}
	if it.l.Size() == 0 {
		it.node = nil
		it.finished = true
		return false
	}

	if it.node == nil { // first time
		it.node = findMin(it.l.root)
	} else {
		cur := it.node
		if cur.r != nil {
			cur = findMin(cur.r)
		} else {
			for {
				if cur.p == nil {
					it.node = nil
					it.finished = true
					return false
				}
				p := cur.p
				pcur := cur
				cur = p
				if cur.l == pcur {
					break
				}
			}
		}
		it.node = cur
	}
	return true
}

func (it *Iterator) Value() interface{} {
	return it.node.val
}
