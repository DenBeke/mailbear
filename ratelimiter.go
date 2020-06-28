package mailbear

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// RateLimitMiddleware copied from https://gitter.im/labstack/echo?at=5a90e681a2194eb80da6faff
func RateLimitMiddleware(period time.Duration, limit int64) echo.MiddlewareFunc {
	rate := limiter.Rate{Period: period, Limit: limit}
	store := memory.NewStore()
	limiterInstance := limiter.New(store, rate)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			context, err := limiterInstance.Get(c.Request().Context(), c.RealIP())
			if err != nil {
				//fmt.Errorf("Could not get context for IP %s - %v", c.RealIP(), err)
				return next(c)
			}
			c.Response().Header().Add("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
			c.Response().Header().Add("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
			c.Response().Header().Add("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))
			if context.Reached {
				return c.JSON(429, map[string]string{"message": fmt.Sprintf("Rate limit exceeded for %s", c.RealIP())})
			}
			return next(c)
		}
	}
}
