package generics

func ContainsMap[K, V comparable](m map[K]V, elem V) bool {
	for _, mapElem := range m {
		if mapElem == elem {
			return true
		}
	}

	return false
}
