package types

type User struct {
	Id      string  `json:"id" db:"id" validate:"required"`
	Name    string  `json:"name" db:"name" validate:"required,max=100"`
	Balance float64 `json:"balance" db:"balance" validate:"required"`
}

func ValidateUser(user *User) bool {
	return true
}
