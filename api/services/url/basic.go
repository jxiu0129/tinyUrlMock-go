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
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

func (s *Service) DeleteByShortenUrl(url *edb.Url) error {
	return s.db.Delete(&edb.Url{}).Where("ShortenUrl = ?", url.ShortenUrl).Error
}
