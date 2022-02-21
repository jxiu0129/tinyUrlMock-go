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
			fmt.Printf("err ==> %+v\n", err)
			return nil, err
		}
		usedKeys = append(usedKeys, k)
	}

	// fmt.Printf("unUsedKeys ==> %+v\n", usedKeys)

	return usedKeys, nil
}
