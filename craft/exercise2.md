# Fix the upper layer

Let's now fix the http layers to use the new implementation.

In [Server structure](./pkg/http/server.go):

1. Remove the `db *sql.DB` field from the Server struct
2. Add a new field `user user.Lister` in the structure
3. Change the `NewServer` signature to accept a `user.Lister` parameter instead of `*sql.DB` and assign it to the Server structure

In [HTTP handler](./pkg/http/get.go):

1. Modify the call `user.List(s.db)` to use `s.user.List()` method instead.