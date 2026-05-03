package gosonarqualitygate

import "cmp"

func Min[T cmp.Ordered](values ...T) (result T, ok bool) {
	if len(values) == 0 {
		return
	}

	result, ok = values[0], true

	for i := 1; i < len(values); i++ {
		if values[i] < result {
			result = values[i]
		}
	}
	return
}
