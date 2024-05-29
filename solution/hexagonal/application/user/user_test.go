package user

import "testing"

func TestValidateUser(t *testing.T) {
	// Arrange
	user := &User{
		Firstname: "Jane",
		Lastname:  "Doe",
		Email:     "jane@doe.com",
	}

	// Act
	err := user.validate()

	// Assert
	if err != nil {
		t.Fatalf("unexpected error while validating user: %s", err)
	}
}

func TestValidateUserInvalidEmail(t *testing.T) {
	// Arrange
	user := &User{
		Firstname: "Jane",
		Lastname:  "Doe",
		Email:     "foobar",
	}

	// Act
	err := user.validate()

	// Assert
	if err == nil {
		t.Fatalf("expecting error while validating user, got nil")
	}
}
