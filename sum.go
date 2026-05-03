package gosonarqualitygate

import "cmp"

func Sum[E cmp.Ordered](values ...E) (total E) {
	for _, v := range values {
		total += v
	}
	return
}
