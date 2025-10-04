package userrepo

import "backend/internal/domain/user"

type InMemUserRepo struct {
	users []user.User
}

func (imur *InMemUserRepo) Save(u *user.User) *user.User {
	imur.users = append(imur.users, *u)
	return u
}

func New() InMemUserRepo {
	return InMemUserRepo{users: []user.User{}}
}
