package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

// Service interface
type Service interface {
	GenerateToken(UserID int) (string, error)
}

type jwtService struct {
}

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// NewService function
func NewService() *jwtService {
	return &jwtService{}
}

// GenerateToken function
func (s *jwtService) GenerateToken(UserID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = UserID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(secretKey)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
