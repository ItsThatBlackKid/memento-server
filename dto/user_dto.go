package dto

type UserDTO struct {
	ID        int8   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
