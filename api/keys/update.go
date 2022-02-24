package keys

import (
	"fmt"
	skey "tinyUrlMock-go/api/services/key"
	"tinyUrlMock-go/lib/db"
)

func SetOneKeyUsed() (string, error) {

	var err error
	// 1. find One From UnusedKey
	key, err := skey.New(db.DBGorm).FindOneUnusedKey()
	if err != nil {
		return "", err
	}

	fmt.Println(key, err)
	updateKeys := []string{key}
	// 2. delete One From UnusedKey
	if err := skey.New(db.DBGorm).DeleteUnusedKeys(updateKeys); err != nil {
		return "", err
	}
	// 3. insertUsedKey
	if err := skey.New(db.DBGorm).InsertUsedKeys(updateKeys); err != nil {
		return "", err
	}

	return key, nil
}

func SetKeyUnused(keys []string) error {

	// keys => ['ws231w', 'dqwdw2',...]
	// var err error
	if err := skey.New(db.DBGorm).DeleteUsedKeys(keys); err != nil {
		return err
	}
	if err := skey.New(db.DBGorm).InsertUnusedKeys(keys); err != nil {
		return err
	}
	return nil
}
