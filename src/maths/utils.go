package maths

func Abs[T Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
