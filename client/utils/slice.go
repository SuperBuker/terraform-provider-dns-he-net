package utils

func ApplyToSlice[T any](fn func(T) T, slice []T) []T {
	for i, v := range slice {
		slice[i] = fn(v)
	}

	return slice
}
