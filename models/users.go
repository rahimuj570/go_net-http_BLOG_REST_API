package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` //we used "-" for ommit this Password in json output
	Role     string `json:"role"`
}
