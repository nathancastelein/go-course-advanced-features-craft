package generics

import "testing"

func TestQueue(t *testing.T) {
	// Arrange
	queue := &Queue[int]{}

	// Act
	for i := 1; i <= 10; i++ {
		queue.Push(i)
	}

	for i := 1; i <= 10; i++ {
		element, err := queue.Pop()

		// Assert
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}
		if element != i {
			t.Fatalf("expected element %d, got %d", i, element)
		}
	}
}
