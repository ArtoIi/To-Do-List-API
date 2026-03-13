package security

import (
	"fmt"
	"os"
	"time"

	"github.com/ArtoIi/To-Do-List-API/internal/domain"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(bytes), err

}

func GenerateToken(u *domain.User) (string, error) {

	secret := os.Getenv("JWT_KEY")

	claims := jwt.MapClaims{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckPassword(password, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {

	key := os.Getenv("JWT_KEY")
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado")
		}
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return *claims, nil
	}

	return nil, fmt.Errorf("token inválido")
}
