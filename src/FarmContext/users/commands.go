package users

type CreateUserCommand struct {
	Username string
	Password string
	FarmID   uint
	FirstName string
	LastName string
	Email string
	Phone string
}

