package dto

import (
	"net/http"
	"strings"

	"github.com/rahimuj570/go_net-http_BLOG_REST_API/enums"
	"github.com/rahimuj570/go_net-http_BLOG_REST_API/models"
)

// DTO For user registrartion
type RegisterUserDTO struct {
	Name     string     `json:"user_name"`
	Email    string     `json:"user_email"`
	Password string     `json:"-"`
	Role     enums.Role `json:"user_role"`
}

// RegisterUserDTO to model.User Convert
func (d *RegisterUserDTO) ToModel() models.User {
	return models.User{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
		Role:     d.Role,
	}
}

// model.User to RegisterUserDTO Convert
func FromModel(m models.User) RegisterUserDTO {
	return RegisterUserDTO{
		Name:     m.Name,
		Email:    m.Email,
		Password: m.Password,
		Role:     m.Role,
	}
}

// RegisterUserDTO validation
func (d *RegisterUserDTO) Validation(w http.ResponseWriter) {
	if len(strings.Split(d.Email, "@")) != 2 {
		http.Error(w, "Email is not valid", http.StatusUnprocessableEntity)
		return
	}
	if len(strings.Split(strings.Split(d.Email, "@")[1], ".")) != 2 {
		http.Error(w, "Email is not valid", http.StatusUnprocessableEntity)
		return
	}
	if d.Name == "" {
		http.Error(w, "Invalid Name", http.StatusUnprocessableEntity)
		return
	}
	if d.Password == "" || len(d.Password) < 4 {
		http.Error(w, "Invalid Password. Must be atleast 4 char", http.StatusUnprocessableEntity)
		return
	}
	if d.Role != enums.READER && d.Role != enums.AUTHOR {
		http.Error(w, "New user must be Reader or Author", http.StatusBadRequest)
		return
	}
}

// DTO For login
type LoginUserDTO struct {
	Email    string `json:"user_email"`
	Password string `json:"-"`
}

// LoginUserDTO validation
func (d *LoginUserDTO) Validation(w http.ResponseWriter) {
	if len(strings.Split(d.Email, "@")) != 2 {
		http.Error(w, "Email is not valid", http.StatusUnprocessableEntity)
		return
	}
	if len(strings.Split(strings.Split(d.Email, "@")[1], ".")) != 2 {
		http.Error(w, "Email is not valid", http.StatusUnprocessableEntity)
		return
	}
	if d.Password == "" || len(d.Password) < 4 {
		http.Error(w, "Invalid Password. Must be atleast 4 char", http.StatusUnprocessableEntity)
		return
	}
}

// DTO For User response
type UserResponseDTO struct {
	ID    int        `json:"user_id"`
	Name  string     `json:"user_name"`
	Email string     `json:"user_email"`
	Role  enums.Role `json:"user_role"`
}
