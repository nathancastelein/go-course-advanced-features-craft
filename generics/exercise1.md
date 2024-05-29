# Exercise 1

## Step 1

Open [./contains.go](contains.go) file.

Write the body of the function to return true if the string `elem` is contained in `slice`.

Test your function:

```go
go test -run TestContainsString .
```

## Step 2

Now we want the same function to check if an integer is contained into a slice.

Write a new function: 
```go
func ContainsInteger(slice []int, elem int) bool {}
```

Don't forget to write the proper unit test in [contains_test.go](./contains_test.go).