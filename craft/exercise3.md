# Fix unit test

At this point, unit test is not working anymore due to our changes.

Let's fix the tests!

## A stub for user.Lister

For this unit test, we will use a stub for the interface `user.Lister`.

1. Add a new `type UserListerStub struct {}`
2. Implement the `user.Lister` interface on this structure, by adding the `List` method

The `List` should return static data:

```go
    return []user.User{
		{
			Firstname: "SpongeBob",
			Lastname:  "SquarePants",
		},
		{
			Firstname: "Patrick",
			Lastname:  "Star",
		},
	}, nil
```

## Use this stub

Fix the Arrange section of the unit test:

1. Remove everything linked to SQL mocks
2. Set a new `UserListerStub` for the `user.Lister` field of `Server` type

## Run test

```
go test -run TestGetUser
```

It should work!

## Last but not least: main function

Finally, we will fix our main function.

Open [main.go](./cmd/api/main.go).

Replace `server := http.NewServer(8080, db)` with `server := http.NewServer(8080, adapters.NewUserSQL(db))`.

Well done!