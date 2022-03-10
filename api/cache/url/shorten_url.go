package ucache

import (
	"encoding/json"
	"time"
	"tinyUrlMock-go/api/entities/edb"
	rlib "tinyUrlMock-go/lib/redis"

	"github.com/gomodule/redigo/redis"
)

// map[originalUrl]{uniqueKey}

type ShortenUrlCache struct {
	ShortenUrl string    `json: shortenUrl`
	CreatedAt  time.Time `json: createdAt`
}

func SetUniqueKey(url *edb.Url) error {

	r := rlib.Pool.Get()
	defer r.Close()

	s := ShortenUrlCache{ShortenUrl: url.ShortenUrl, CreatedAt: url.CreatedAt}

	sc, err := json.Marshal(&s)
	if err != nil {
		return err
	}

	if _, err := r.Do(rlib.SET, url.OriginalUrl, string(sc)); err != nil {
		return err
	}

	return nil
}

func GetUniqueKey(originalUrl string) (ShortenUrlCache, error) {

	r := rlib.Pool.Get()
	defer r.Close()
	urlStr, err := redis.String(r.Do(rlib.GET, originalUrl))
	if err != nil {
		if err.Error() == redis.ErrNil.Error() {
			return ShortenUrlCache{}, nil
		}
	}

	jsonUrl := ShortenUrlCache{}
	json.Unmarshal([]byte(urlStr), &jsonUrl)

	if err != nil {
		return ShortenUrlCache{}, err
	}
	return jsonUrl, nil
}
