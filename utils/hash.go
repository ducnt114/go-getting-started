package utils

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
)

const (
	// Pbkdf2Iterations sets the amount of iterations used by the PBKDF2 hashing algorithm
	Pbkdf2Iterations int = 15000
	// HashBytes sets the amount of bytes for the hash output from the PBKDF2 / scrypt hashing algorithm
	HashBytes int = 64
	// UniqueKey Key
	UniqueKey = "!!!!"
)

func HashPassword(rawPass, saltPassword string) string {
	key, _ := base64.StdEncoding.DecodeString(saltPassword)
	bytes := pbkdf2.Key([]byte(rawPass+UniqueKey), key, Pbkdf2Iterations, HashBytes, sha512.New)
	return hex.EncodeToString(bytes)
}
