package surl

import (
	"fmt"
	"tinyUrlMock-go/api/entities/edb"
)

func (s *Service) InsertUrls(newUrlsArray []*edb.Url) error {
	valuesStr := ""
	for i, key := range newUrlsArray {
		if i == len(newUrlsArray)-1 {
			valuesStr += fmt.Sprintf("('%v','%v','%v')", key.ShortenUrl, key.OriginalUrl, key.CreatedAt)
			break
		}
		valuesStr += fmt.Sprintf("('%v','%v','%v'),", key.ShortenUrl, key.OriginalUrl, key.CreatedAt)
	}

	sql := "INSERT INTO" +
		"`Url` (`ShortenUrl`, `OriginalUrl`, `CreatedAt`)" +
		"VALUES " + valuesStr

	if err := s.db.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) FindExistUrl(url string) (*edb.Url, error) {
	u := &edb.Url{}
	if err := s.db.Where("OriginalUrl = ?", url).First(u).Error; err != nil {
		if err.Error() == "record not found" {
			// 若db查無此筆
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}
