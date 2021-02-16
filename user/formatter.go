package user

// UserFormatter struct
type UserFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Ocupation string `json:"ocupation"`
	Avatar    string `json:"avatar"`
	Role      string `json:"role"`
	Token     string `json:"token"`
}

// FormatUser function
func FormatUser(user User, token string) UserFormatter {
	FormattedUser := UserFormatter{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Ocupation: user.Ocupation,
		Avatar:    user.Avatar,
		Role:      user.Role,
		Token:     token,
	}

	return FormattedUser
}
