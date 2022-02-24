package url

import (
	"fmt"
	"tinyUrlMock-go/api/entities/edb"
	"tinyUrlMock-go/api/keys"
	surl "tinyUrlMock-go/api/services/url"
	"tinyUrlMock-go/lib/db"
)

func UrlExpired(url *edb.Url) error {
	fmt.Println("url expired")
	if err := surl.New(db.DBGorm).DeleteByShortenUrl(url); err != nil {
		return err
	}
	if err := keys.SetKeyUnused([]string{url.ShortenUrl}); err != nil {
		return err
	}
	return nil
}
