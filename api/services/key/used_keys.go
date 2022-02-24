package skey

import (
	"fmt"
	"tinyUrlMock-go/api/entities/edb"
)

func (s *Service) SearchAllUsedKeys() ([]*edb.UsedKeys, error) {
	// allKeys := &[]edb.UsedKeys{}
	usedKey := &(edb.UsedKeys{})
	usedKeys := []*edb.UsedKeys{}

	usedKeyRows, err := s.db.Model(usedKey).
		Select("UniqueKey").
		Rows()

	if err != nil {
		return nil, err
	}
	defer usedKeyRows.Close()

	// fmt.Printf("rows ==> %+v\n", rows)
	for usedKeyRows.Next() {
		k := &edb.UsedKeys{}
		if err := s.db.ScanRows(usedKeyRows, k); err != nil {
			return nil, err
		}
		usedKeys = append(usedKeys, k)
	}

	// fmt.Printf("unUsedKeys ==> %+v\n", usedKeys)

	return usedKeys, nil
}

func (s *Service) InsertUsedKeys(newKeysArray []string) error {
	valuesStr := ""
	for i, key := range newKeysArray {
		if i == len(newKeysArray)-1 {
			valuesStr += fmt.Sprintf("('%v')", key)
			break
		}
		valuesStr += fmt.Sprintf("('%v'),", key)
	}

	sql := "INSERT INTO" +
		"`UsedKeys` (`UniqueKey`)" +
		"VALUES " + valuesStr

	if err := s.db.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteUsedKeys(keysArray []string) error {
	return s.db.Delete(&edb.UsedKeys{}, "UniqueKey IN (?)", keysArray).Error
}
