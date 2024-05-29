package generics

import "testing"

func TestCountString(t *testing.T) {
	// Arrange
	slice := []string{"Hello", "World", "Hello"}
	countElem := "Hello"

	// Act
	count := Count(slice, countElem)

	// Assert
	if count != 2 {
		t.Fatalf("expected count to be 2, got %d", count)
	}
}

func TestCountInt(t *testing.T) {
	// Arrange
	slice := []int{1, 2, 3, 4, 3, 2, 1}
	countElem := 2

	// Act
	count := Count(slice, countElem)

	// Assert
	if count != 2 {
		t.Fatalf("expected count to be 2, got %d", count)
	}
}

func TestCountIntSpecified(t *testing.T) {
	// Arrange
	slice := []int{1, 2, 3, 4, 3, 2, 1}
	countElem := 2

	// Act
	count := Count[int](slice, countElem)

	// Assert
	if count != 2 {
		t.Fatalf("expected count to be 2, got %d", count)
	}
}

func TestCountMap(t *testing.T) {
	// Arrange
	m := map[string]int{
		"Hello": 1,
		"World": 2,
		"!":     2,
	}
	countElem := 2

	// Act
	count := CountMap(m, countElem)

	// Assert
	if count != 2 {
		t.Fatalf("expected count to be 2, got %d", count)
	}
}
