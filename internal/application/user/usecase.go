package usercase

import (
	"backend/internal/domain/user"
	"backend/internal/infrastructure/persistance/userrepo"
)

type RegisterUserUseCase struct {
	userRepo user.Repository
}

func NewRegisterUserUseCase() *RegisterUserUseCase {
	userRepo := userrepo.New()
	return &RegisterUserUseCase{userRepo: &userRepo}
}

func (uc *RegisterUserUseCase) Execute(email, name, password string) (*user.User, error) {
	u, err := user.New(email, name, password)
	if err != nil {
		return &user.User{}, err
	}

	return uc.userRepo.Save(u), nil
}
