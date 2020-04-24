package auth

import (
	"bytes"
	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) []byte {
	bKey := argon2.Key(
		[]byte(password),
		[]byte("secret"),
		3, 32*1024, 4, 32)
	return bKey
}

func CheckPasswordHash(password string, hash []byte) bool {
	return bytes.Equal(
		hash,
		HashPassword(password),
	)
}
