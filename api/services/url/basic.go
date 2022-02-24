package surl

import (
	"fmt"
	"tinyUrlMock-go/api/entities/edb"

	"github.com/jinzhu/gorm"
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

type FindUrl struct {
	OriginalUrl string
	ShortenUrl  string
}

// ! will delete, leave for check
/* func (s *Service) FindExistUrl(url FindUrl) (*edb.Url, error) {
	u := &edb.Url{}
	if url.ShortenUrl != "" {
		if err := s.db.Where("ShortenUrl = ?", url.ShortenUrl).First(u).Error; err != nil {
			if err.Error() == "record not found" {
				// 若db查無此筆
				return nil, nil
			}
			return nil, err
		}
		return u, nil
	}
	if err := s.db.Where("OriginalUrl = ?", url.OriginalUrl).First(u).Error; err != nil {
		if err.Error() == "record not found" {
			// 若db查無此筆
			return nil, nil
		}
		return nil, err
	}
	return u, nil
} */

func (s *Service) FindShortenUrl(url string) (*edb.Url, error) {
	u := &edb.Url{}
	if err := s.db.Where("ShortenUrl = ?", url).First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
func (s *Service) FindOriginalUrl(url string) (*edb.Url, error) {
	u := &edb.Url{}
	if err := s.db.Where("OriginalUrl = ?", url).First(u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// 若db查無此筆
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

func (s *Service) DeleteByShortenUrl(url *edb.Url) error {
	return s.db.Delete(&edb.Url{}).Where("ShortenUrl = ?", url.ShortenUrl).Error
}
