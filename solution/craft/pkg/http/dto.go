package http

import "github.com/nathancastelein/go-course-advanced-craft/advanced-features/solution/craft/pkg/user"

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (u *User) From(domainUser user.User) User {
	return User{
		Firstname: domainUser.Firstname,
		Lastname:  domainUser.Lastname,
	}
}
