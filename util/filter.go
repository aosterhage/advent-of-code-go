package util

func Filter[T any](ts []T, test func(T) bool) (ret []T) {
	for _, t := range ts {
		if test(t) {
			ret = append(ret, t)
		}
	}
	return
}
