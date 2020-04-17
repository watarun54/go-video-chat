package domain

type Users []User

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}
