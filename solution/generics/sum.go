package generics

func Sum[T int64 | float64](slice []T) T {
	var total T
	for _, element := range slice {
		total += element
	}

	return total
}
