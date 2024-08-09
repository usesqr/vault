package db

import (
	"fmt"

	sqliteEncrypt "github.com/jackfr0st13/gorm-sqlite-cipher"
	"github.com/usesqr/vault/crypto"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(password string, isFirstTime bool) {
	var keyStr string
	if isFirstTime {
		keyStr = string(crypto.GenerateNewCryptoKey(password))
	} else {
		keyStr = string(crypto.DeriveExistingCryptoKey(password))
	}
	dbNameWithDSN := GetDatabaseFilePath() + fmt.Sprintf("?_pragma_key=%s&_pragma_cipher_page_size=4096", keyStr)
	db, err := gorm.Open(sqliteEncrypt.Open(dbNameWithDSN))

	if err != nil {
		panic(err)
	}

	DB = db
}
