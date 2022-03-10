package keys

import (
	"fmt"
	skey "tinyUrlMock-go/api/services/key"
	"tinyUrlMock-go/lib/db"
)

func SetOneKeyUsed() (string, error) {

	var err error

	key, err := skey.New(db.DBGorm).FindOneUnusedKey()
	if err != nil {
		return "", err
	}

	fmt.Println(key, err)
	updateKeys := []string{key}

	if err := skey.New(db.DBGorm).DeleteUnusedKeys(updateKeys); err != nil {
		return "", err
	}

	if err := skey.New(db.DBGorm).InsertUsedKeys(updateKeys); err != nil {
		return "", err
	}

	return key, nil
}

func SetKeyUnused(keys []string) error {

	if err := skey.New(db.DBGorm).DeleteUsedKeys(keys); err != nil {
		return err
	}
	if err := skey.New(db.DBGorm).InsertUnusedKeys(keys); err != nil {
		return err
	}
	return nil
}
