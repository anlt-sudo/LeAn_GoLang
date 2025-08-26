package user

type User struct {
	Email string
	Name  string
}

func NewUser(email, name string) *User {
	return &User{Email: email, Name: name}
}
