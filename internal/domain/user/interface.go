package user

type Repository interface {
	Save(u *User) *User
}
