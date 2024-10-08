package db

import (
	"encoding/base64"
	"fmt"

	sqliteEncrypt "github.com/jackfr0st13/gorm-sqlite-cipher"
	"github.com/usesqr/vault/crypto"
	"github.com/usesqr/vault/db/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(password string, isFirstTime bool) error {
	var key []byte
	if isFirstTime {
		key = crypto.GenerateNewCryptoKey(password)
	} else {
		key = crypto.DeriveExistingCryptoKey(password)
	}
	keyStr := base64.StdEncoding.EncodeToString(key)
	dbNameWithDSN := GetDatabaseFilePath() + fmt.Sprintf("?_pragma_key=%s&_pragma_cipher_page_size=4096", keyStr)
	db, err := gorm.Open(sqliteEncrypt.Open(dbNameWithDSN))

	if err != nil {
		return err
	}

	db.AutoMigrate(&model.Password{})

	DB = db
	return nil
}
