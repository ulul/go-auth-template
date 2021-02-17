package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// Service interface
type Service interface {
	GenerateToken(UserID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
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

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, valid := t.Method.(*jwt.SigningMethodHMAC)

		if !valid {
			return nil, errors.New("Invalid token")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
