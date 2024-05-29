package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type RepositoryStub struct {
	users              []User
	emailAlreadyExists bool
	errorOnSave        bool
}

func (r *RepositoryStub) List(ctx context.Context) ([]User, error) {
	return r.users, nil
}

func (r *RepositoryStub) EmailAlreadyExists(ctx context.Context, email string) (bool, error) {
	return r.emailAlreadyExists, nil
}

func (r *RepositoryStub) Save(ctx context.Context, user User) error {
	if r.errorOnSave {
		return errors.New("something wrong happened")
	}
	return nil
}

func TestList(t *testing.T) {
	// Arrange
	repositoryStub := &RepositoryStub{
		users: []User{
			{
				Firstname: "Jane",
				Lastname:  "Doe",
				Email:     "jane@doe.com",
			},
		},
		emailAlreadyExists: false,
		errorOnSave:        false,
	}
	service := NewService(repositoryStub)

	// Act
	users, err := service.List(context.Background())

	// Assert
	require.NoError(t, err)
	require.Len(t, users, 1)
}

func TestCreate(t *testing.T) {
	// Arrange
	repositoryStub := &RepositoryStub{
		users:              []User{},
		emailAlreadyExists: false,
		errorOnSave:        false,
	}
	service := NewService(repositoryStub)
	expectedUser := User{
		Firstname: "Jane",
		Lastname:  "Doe",
		Email:     "jane@doe.com",
	}

	// Act
	user, err := service.Create(context.Background(), "Jane", "Doe", "jane@doe.com")

	// Assert
	require.NoError(t, err)
	require.Equal(t, expectedUser, user)
}

func TestCreateUserAlreadyExists(t *testing.T) {
	// Arrange
	repositoryStub := &RepositoryStub{
		users:              []User{},
		emailAlreadyExists: true,
		errorOnSave:        false,
	}
	service := NewService(repositoryStub)
	expectedUser := User{}

	// Act
	user, err := service.Create(context.Background(), "Jane", "Doe", "jane@doe.com")

	// Assert
	require.Error(t, err)
	require.Equal(t, expectedUser, user)
}

func TestCreateUserErrorOnSave(t *testing.T) {
	// Arrange
	repositoryStub := &RepositoryStub{
		users:              []User{},
		emailAlreadyExists: false,
		errorOnSave:        true,
	}
	service := NewService(repositoryStub)
	expectedUser := User{}

	// Act
	user, err := service.Create(context.Background(), "Jane", "Doe", "jane@doe.com")

	// Assert
	require.Error(t, err)
	require.Equal(t, expectedUser, user)
}
