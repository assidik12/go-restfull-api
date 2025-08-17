package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/assidik12/go-restfull-api/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims adalah struct untuk data yang ingin kita simpan di dalam token.
type CustomClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

// GenerateJWT membuat token baru untuk user.
// Ia menerima ID dan email user sebagai input untuk dimasukkan ke dalam claims.
func GenerateJWT(user domain.User) (string, error) {
	// Ambil secret key dari environment variable.
	// PASTIKAN ANDA SUDAH MENETAPKAN JWT_SECRET DI FILE .env ANDA.
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("JWT_SECRET environment variable not set")
	}

	claims := CustomClaims{
		UserID: user.ID,
		Email:  user.Email,
		Name:   user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// Buat token dengan claims dan metode signing HS256.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token dengan secret key Anda.
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateJWT memvalidasi token yang diberikan dan mengembalikan claims jika valid.
func ValidateJWT(tokenString string) (*CustomClaims, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET environment variable not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Jika token valid, kita bisa mengakses claims-nya.
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
