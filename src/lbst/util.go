package lbst

type Comparator func(a, b interface{}) int

func IntComparator(a, b interface{}) int {
	aa := a.(int)
	bb := b.(int)
	switch {
	case aa > bb:
		return 1
	case aa < bb:
		return -1
	default:
		return 0
	}
}
