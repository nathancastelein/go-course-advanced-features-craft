package user

import (
	"errors"
	"regexp"
)

type User struct {
	Firstname string
	Lastname  string
	Email     string
}

func (u User) validate() error {
	if u.Firstname == "" {
		return errors.New("invalid firstname, should not be empty")
	}

	if u.Lastname == "" {
		return errors.New("invalid lastname, should not be empty")
	}

	if u.Email == "" {
		return errors.New("invalid email, should not be empty")
	}

	if !regexp.MustCompile(`^.*@.*\..*$`).MatchString(u.Email) {
		return errors.New("invalid email, invalid format")
	}

	return nil
}
