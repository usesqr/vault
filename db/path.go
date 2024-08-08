package db

import (
	"os"
	"path"

	"github.com/usesqr/vault/consts"
)

func GetDatabaseFilePath() string {
	confDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	return path.Join(confDir, consts.Identifier, "vault_data.db")
}