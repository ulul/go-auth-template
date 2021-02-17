package user

// RegisterUserInput struct
type RegisterUserInput struct {
	Name      string `json:"name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Ocupation string `json:"ocupation" binding:"required"`
}

// LoginInput struct
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
