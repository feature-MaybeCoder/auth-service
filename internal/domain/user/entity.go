package user

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password Password
}

func processPassword(password string) (Password, error) {

	if password == "" {
		return Password{Value: "", IsSet: false}, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, err
	}
	return Password{Value: string(hashedPassword), IsSet: true}, nil
}

func New(email, name string, password string) (*User, error) {
	uuid := uuid.New()

	processedPassword, err := processPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{ID: uuid, Name: name, Email: email, Password: processedPassword}, nil
}
