package crypto

import (
	"crypto/rand"
	"errors"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/usesqr/vault/consts"
)

func GetSaltFilePath() string {
	confDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	return path.Join(confDir, consts.Identifier, "salt.bin")
}

func GenerateSalt() []byte {
	salt := make([]byte, 8)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		panic(err)
	}

	return salt
}

func WriteSaltFile(salt []byte) {
	path := GetSaltFilePath()

	// Ensure the directory exists
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		panic(err)
	}

	// Write the salt to the file
	err = os.WriteFile(path, salt, 0644)
	if err != nil {
		panic(err)
	}
}

func ReadSaltFile() []byte {
	data, err := os.ReadFile(GetSaltFilePath())
	if err != nil {
		panic(err)
	}

	return data
}

func DoesSaltFileExist() bool {
	_, err := os.ReadFile(GetSaltFilePath())
	if err != nil {
		return !errors.Is(err, fs.ErrNotExist)
	}
	return true
}
