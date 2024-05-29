package generics

import "testing"

func TestContainsString(t *testing.T) {
	// Arrange
	slice := []string{"Hello", "World"}
	elem := "Hello"

	// Act
	contains := ContainsString(slice, elem)

	// Assert
	if !contains {
		t.Fatalf("%s expected to be contained in %s, got false", elem, slice)
	}
}
