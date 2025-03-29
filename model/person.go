package model

type Person struct {
	ID        int    `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	Phone     string `json:"phone" db:"phone"`
	FirstName string `json:"firstname" db:"firstname"`
	LastName  string `json:"lastname" db:"lastname"`
}
