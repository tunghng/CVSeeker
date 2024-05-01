package utils

import "golang.org/x/crypto/bcrypt"

// IsValidPassword verifies whether a password is the user password
func IsValidPassword(password string, token string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(token), []byte(password))
	return err == nil
}

// SetPassword encrypt password, and set to EncryptedPassword field
func SetPassword(password string) string {
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(encryptedPassword)
}
