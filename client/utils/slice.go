package utils

func ApplyToSlice[T any](fn func(T) T, slice []T) []T {
	for i, v := range slice {
		v := v // To deprecate in Go v1.22
		slice[i] = fn(v)
	}

	return slice
}
