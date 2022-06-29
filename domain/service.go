package domain

type Service interface {
	Create(user *User) error
}

type Repository interface {
	Create(user *User) error
}
