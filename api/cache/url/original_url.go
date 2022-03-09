package ucache

import (
	"encoding/json"
	"fmt"
	"time"
	"tinyUrlMock-go/api/entities/edb"
	rlib "tinyUrlMock-go/lib/redis"

	"github.com/gomodule/redigo/redis"
)

// map[shortenUrl]{OriginalUrl, createat}

type OriginalUrlCache struct {
	OriginalUrl string    `json: originalUrl`
	CreatedAt   time.Time `json: createdAt`
}

func SetOriginalUrl(url *edb.Url) error {

	r := rlib.Pool.Get()
	defer r.Close()

	o := OriginalUrlCache{OriginalUrl: url.OriginalUrl, CreatedAt: url.CreatedAt}

	oc, err := json.Marshal(&o)
	if err != nil {
		return err
	}

	if _, err := r.Do(rlib.SET, url.ShortenUrl, string(oc)); err != nil {
		return err
	}

	return nil
}

func GetOriginalUrl(shortenUrl string) (OriginalUrlCache, error) {

	r := rlib.Pool.Get()
	defer r.Close()

	urlStr, err := redis.String(r.Do(rlib.GET, shortenUrl))
	fmt.Println(urlStr)

	if err != nil {
		if err.Error() == redis.ErrNil.Error() {
			return OriginalUrlCache{}, nil
		}
	}

	jsonUrl := OriginalUrlCache{}
	json.Unmarshal([]byte(urlStr), &jsonUrl)

	if err != nil {
		return OriginalUrlCache{}, err
	}
	return jsonUrl, nil
}
