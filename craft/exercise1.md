# A SOLID API

Start by having a look on the current code.

This project is an API, to list users from a database.

Have a look on the different files of the project.

Let's now rewrite a bit of stuff to make this code more in adequation with SOLID principles, and more specifically the dependency inversion principle.

## The issue

Have a look on the [unit test of the http handler](./pkg/http/get_test.go). Is there something wrong with this test?

The highest layer of the application (the HTTP handler) depends directly on the lowest layer of it (the SQL data storage).

So all layers are currently strongly coupled, and we want to avoid this situation.

## Step 1: the user package

For now, the [user package](./pkg/user/) mixes some business objects - `User` structure - with some data storage mecanisms (finding data from the database).

Consider the `user` package as the core package/the business package of the application.

Let's remove SQL from this package!

1. Create a new package `adapters` (we will discuss this name after) in `/pkg`.
2. Create a new struct `UserSQL` in this package. This struct contains one field: `db *sql.DB`
3. Create a constructor function `NewUserSQL(db *sql.DB) *UserSQL` to return a `UserSQL` properly created.
4. Move the `List()` function and attach it as a method for `UserSQL`: `func (u *UserSQL) List() ([]user.User, error)`. The input parameter `db` is no longer required, use the one from the `UserSQL` structure!
5. Move the [test file](./pkg/user/user_test.go) to the `adapters` package. Fix the test and launch it!

At this point, the `user` package should only contains the `User` struct.
The `adapters` package contains two files, and `go test` is working on this package.

The others packages `http` and `main` are broken, it's OK for now.

## Step 2: a new interface in user

As we discussed, interfaces are a viable option to inverse dependencies. Let's use it!

1. Create a new file in `user` package, named `ports.go`
2. Add a new interface `Lister` with only one method in it: `List() ([]User, error)`

As interface implementation is implicit in Go, the `UserSQL` structure already meets the interface!

Change the signature of the constructor to ensure the interface is implemented.

Instead of `func NewUserSQL(db *sql.DB) *UserSQL`, change it to `func NewUserSQL(db *sql.DB) user.Lister`.

It should just work!