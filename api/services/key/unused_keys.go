package skey

import (
	"fmt"
	"tinyUrlMock-go/api/entities/edb"
)

func (s *Service) SearchAllUnusedKeys() ([]*edb.UnusedKeys, error) {
	unUsedKey := &(edb.UnusedKeys{})
	unUsedKeys := []*edb.UnusedKeys{}

	unUsedKeyRows, err := s.db.Model(unUsedKey).
		Select("UniqueKey").
		Rows()

	if err != nil {
		return nil, err
	}
	defer unUsedKeyRows.Close()

	// fmt.Printf("rows ==> %+v\n", rows)
	for unUsedKeyRows.Next() {
		k := &edb.UnusedKeys{}
		if err := s.db.ScanRows(unUsedKeyRows, k); err != nil {
			fmt.Printf("err ==> %+v\n", err)
			return nil, err
		}
		unUsedKeys = append(unUsedKeys, k)
	}

	return unUsedKeys, nil
}

func (s *Service) InsertNewUnusedKeys(newKeysArray []string) (string, error) {
	valuesStr := ""
	for i, key := range newKeysArray {
		if i == len(newKeysArray)-1 {
			valuesStr += fmt.Sprintf("('%v')", key)
			break
		}
		valuesStr += fmt.Sprintf("('%v'),", key)
	}

	sql := "INSERT INTO" +
		"`UnusedKeys` (`UniqueKey`)" +
		"VALUES " + valuesStr

	if err := s.db.Exec(sql).Error; err != nil {
		return "insertNewKey Error", err
	}

	return "insertNewKey success", nil
}

func (s *Service) FindOneUnusedKey() (string, error) {
	k := &edb.UnusedKeys{}
	if err := s.db.First(k).Error; err != nil {
		return "FindOneUnusedKey error", err
	}
	return k.UniqueKey, nil
}

func (s *Service) DeleteUnusedKeys(keys []string) error {
	return s.db.Delete(&edb.UnusedKeys{}, "UniqueKey IN (?)", keys).Error
}

func (s *Service) InsertUnusedKeys(keys []string) error {
	valuesStr := ""
	for i, key := range keys {
		if i == len(keys)-1 {
			valuesStr += fmt.Sprintf("('%v')", key)
			break
		}
		valuesStr += fmt.Sprintf("('%v'),", key)
	}

	sql := "INSERT INTO" +
		"`UnusedKeys` (`UniqueKey`)" +
		"VALUES " + valuesStr

	if err := s.db.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}
