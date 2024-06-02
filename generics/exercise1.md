# Exercise 2

## Step 1

Write a new function `Contains` with generics, to check if an elem is contained into a slice.

You can add two unit tests to test your generic implementation:

```go
func TestContainsGenericString(t *testing.T) {
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

func TestContainsGenericInt(t *testing.T) {
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
```

Run your tests with:

```bash
go test -run TestContainsGeneric .
```

## Step 2

Open [containsmap.go](./containsmap.go) file.

Rewrite the function `ContainsMap` with generics, to check if `elem` is contained as a value in the given map `m`.
To understand the expected behaviour, you can check the unit test [containsmap_test.go](./containsmap_test.go).

You can then test your implementation with:
```bash
go test -run TestContainsMap .
```

Add this test:

```go
func TestContainsMapString(t *testing.T) {
	// Arrange
	m := map[int]string{
		1: "Bonjour",
		2: "Monde",
	}
	elem := "Monde"

	// Act
	contains := ContainsMap(m, elem)

	// Assert
	if !contains {
		t.Fatalf("%s expected to be contained in %v, got false", elem, m)
	}
}
```

Then run it to confirm your code is using generics:
```bash
go test -run TestContainsMapString .
```