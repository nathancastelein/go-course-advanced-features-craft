# Fix unit test

At this point, unit test is not working anymore due to our changes.

Let's fix the tests!

## A stub for user.Lister

For this unit test, we will use a stub for the interface `user.Lister`.

```go
type UserListerStub struct{}

func (u *UserGetterStub) List() ([]user.User, error) {
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
}
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