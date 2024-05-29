package user

import (
	"context"
	"errors"
)

type Service interface {
	List(context.Context) ([]User, error)
	Create(ctx context.Context, firstname, lastname, email string) (User, error)
}

func NewService(repository Repository) Service {
	return &driverPortImpl{
		repository: repository,
	}
}

type driverPortImpl struct {
	repository Repository
}

func (l *driverPortImpl) List(ctx context.Context) ([]User, error) {
	return l.repository.List(ctx)
}

func (l *driverPortImpl) Create(ctx context.Context, firstname, lastname, email string) (User, error) {
	user := User{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
	}

	if err := user.validate(); err != nil {
		return User{}, err
	}

	userAlreadyExists, err := l.repository.EmailAlreadyExists(ctx, user.Email)
	if err != nil {
		return User{}, err
	}

	if userAlreadyExists {
		return User{}, errors.New("user already exists with this email")
	}

	if err := l.repository.Save(ctx, user); err != nil {
		return User{}, err
	}

	return user, nil
}
