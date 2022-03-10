package middleware

import (
	"net/http"
	"time"
	"tinyUrlMock-go/lib/errors"
	"tinyUrlMock-go/lib/redis"

	"github.com/gin-gonic/gin"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

const (
	DefRateLimiterMaxRetry = 3
	DefRateLimiterPeriod   = 1 * time.Minute
)

func RateLimiterByIP(period time.Duration, limit int64) gin.HandlerFunc {
	return RateLimiter(period, limit)
}

func RateLimiter(period time.Duration, limit int64) gin.HandlerFunc {

	return func(c *gin.Context) {
		rate := limiter.Rate{
			Period: period,
			Limit:  limit,
		}

		store, err := sredis.NewStoreWithOptions(redis.Client, limiter.StoreOptions{
			Prefix:   "rate-limiter",
			MaxRetry: DefRateLimiterMaxRetry,
		})
		if err != nil {
			errors.Error(c, http.StatusInternalServerError, errors.CODE_RATE_LIMITER_ERROR, err)
			return
		}

		middleware := &mgin.Middleware{
			Limiter:        limiter.New(store, rate),
			OnError:        DefaultErrorHandler,
			OnLimitReached: DefaultLimitReachedHandler,
			KeyGetter:      DefaultKeyGetter,
			ExcludedKey:    nil,
		}

		middleware.Handle(c)
	}
}

func DefaultErrorHandler(c *gin.Context, err error) {
	errors.Error(c, http.StatusInternalServerError, errors.CODE_RATE_LIMITER_ERROR, err)
}

func DefaultLimitReachedHandler(c *gin.Context) {
	errors.Throw(c, errors.ErrAuth.Err)
}

func DefaultKeyGetter(c *gin.Context) string {
	return c.ClientIP()
}
