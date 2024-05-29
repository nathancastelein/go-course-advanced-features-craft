package generics

import "testing"

func TestContainsMap(t *testing.T) {
	// Arrange
	m := map[string]int{
		"Hello": 1,
		"World": 2,
	}
	elem := 2

	// Act
	contains := ContainsMap(m, elem)

	// Assert
	if !contains {
		t.Fatalf("%d expected to be contained in %v, got false", elem, m)
	}
}
