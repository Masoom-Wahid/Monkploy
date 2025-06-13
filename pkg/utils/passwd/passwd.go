package passwd

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func HashOTP(otp string) string {
	hashedOTP, err := HashPassword(otp)
	if err != nil {
		return ""
	}

	return string(hashedOTP)
}

func VerifyOTP(hashedCode, code string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedCode), []byte(code))
}
