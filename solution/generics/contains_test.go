package generics

import "testing"

func TestContainsString(t *testing.T) {
	// Arrange
	slice := []string{"Hello", "World"}
	elem := "Hello"

	// Act
	contains := Contains(slice, elem)

	// Assert
	if !contains {
		t.Fatalf("%s expected to be contained in %v, got false", elem, slice)
	}
}

func TestContainsInt(t *testing.T) {
	// Arrange
	slice := []int{1, 2, 3}
	elem := 3

	// Act
	contains := Contains(slice, elem)

	// Assert
	if !contains {
		t.Fatalf("%d expected to be contained in %v, got false", elem, slice)
	}
}
