package user

import "context"

type Repository interface {
	List(context.Context) ([]User, error)
	EmailAlreadyExists(ctx context.Context, email string) (bool, error)
	Save(context.Context, User) error
}
