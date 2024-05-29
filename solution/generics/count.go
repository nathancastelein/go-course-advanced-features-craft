package generics

func Count[T comparable](slice []T, elem T) int {
	count := 0
	for _, sliceElem := range slice {
		if sliceElem == elem {
			count++
		}
	}

	return count
}

func CountMap[K, V comparable](m map[K]V, elem V) int {
	count := 0
	for _, mapElem := range m {
		if mapElem == elem {
			count++
		}
	}

	return count
}
