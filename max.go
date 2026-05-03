package gosonarqualitygate

import "cmp"

func Max[T cmp.Ordered](values ...T) (max T, result bool) {
	if len(values) == 0 {
		return
	}

	for _, val := range values {
		if val > max {
			max = val
		}
	}
	return max, true
}
