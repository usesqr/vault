package crypto

import (
	"crypto/sha1"

	"golang.org/x/crypto/pbkdf2"
)

func GenerateNewCryptoKey(password string) []byte {
	salt := GenerateSalt()
	WriteSaltFile(salt)

	return pbkdf2.Key([]byte(password), salt, 4096, 32, sha1.New)
}

func DeriveExistingCryptoKey(password string) []byte {
	salt := ReadSaltFile()
	
	return pbkdf2.Key([]byte(password), salt, 4096, 32, sha1.New)
}