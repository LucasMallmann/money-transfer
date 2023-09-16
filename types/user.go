package types

type User struct {
	Id    string `json:"id" db:"id" validate:"required"`
	Name  string `json:"nome" db:"nome" validate:"required,max=100"`
	Email string `json:"email" db:"email" validate:"required,email"`
}

func ValidateUser(user *User) bool {
	return true
}
