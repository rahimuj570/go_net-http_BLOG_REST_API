package models

import "github.com/rahimuj570/go_net-http_BLOG_REST_API/enums"

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Role     enums.Role
}
