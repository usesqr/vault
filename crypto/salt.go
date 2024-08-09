package crypto

import (
	"crypto/rand"
	"io"
	"os"
	"path"

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
	err := os.WriteFile(GetSaltFilePath(), salt, 0644)
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
	if err == nil {
		return true
	}

	return os.IsNotExist(err)
}