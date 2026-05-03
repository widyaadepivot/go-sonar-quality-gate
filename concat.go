package gosonarqualitygate

func Concat(values ...string) string {
	var str string
	for _, val := range values {
		str += val
	}
	return str
}
