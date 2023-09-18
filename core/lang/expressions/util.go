package expressions

import "fmt"

func castOrPanic[T any](toCast any) T {
	v, ok := toCast.(T)
	if !ok {
		var e T
		panic(fmt.Sprintf("can't use %T when %T required", toCast, e))
	}
	return v
}
