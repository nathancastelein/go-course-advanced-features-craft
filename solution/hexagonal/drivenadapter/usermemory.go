package drivenadapter

import (
	"context"
	"sync"

	"github.com/nathancastelein/go-course-advanced-craft/advanced-features/solution/hexagonal/application/user"
)

var (
	_ user.Repository = &UserInMemory{}
)

func NewUserInMemory() user.Repository {
	return &UserInMemory{}
}

type UserInMemory struct {
	mu    sync.Mutex
	users map[string]user.User
}

func (u *UserInMemory) List(ctx context.Context) ([]user.User, error) {
	users := make([]user.User, len(u.users))

	var idx int
	for _, user := range users {
		users[idx] = user
		idx++
	}

	return users, nil
}

func (u *UserInMemory) Save(ctx context.Context, user user.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.users[user.Email] = user

	return nil
}

func (u *UserInMemory) EmailAlreadyExists(ctx context.Context, email string) (bool, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	_, ok := u.users[email]
	return ok, nil
}
