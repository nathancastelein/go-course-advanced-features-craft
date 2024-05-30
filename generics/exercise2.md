# Exercise 2

## Step 1

Write a new function `Contains` with generics, to check if an elem is contained into a slice.

Change your tests to use `Contains` instead of `ContainsString` or `ContainsInteger`.

Do not perform other changes in your tests, it should work!

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