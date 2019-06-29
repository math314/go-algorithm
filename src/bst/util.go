package bst

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

func Max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

type ComparableInt int

func (p *ComparableInt) Compare(r interface{}) CompareResult {
	rval, _ := r.(*ComparableInt)
	return ToCompareResult(int(*p) - int(*rval))
}
