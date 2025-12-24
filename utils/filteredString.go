package utils

func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, element := range slice {
		if predicate(element) {
			result = append(result, element)
		}
	}
	return result
}
