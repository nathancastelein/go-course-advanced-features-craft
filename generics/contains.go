package generics

func ContainsString(slice []string, elem string) bool {
	for _, sliceElem := range slice {
		if sliceElem == elem {
			return true
		}
	}

	return false
}

func ContainsInt(slice []int, elem int) bool {
	for _, sliceElem := range slice {
		if sliceElem == elem {
			return true
		}
	}

	return false
}
