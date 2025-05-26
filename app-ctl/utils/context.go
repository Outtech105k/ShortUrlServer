package utils

import (
	redisclient "github.com/Outtech105k/ShortUrlServer/app-ctl/redis-client"
)

type AppContext struct {
	Config Config
	Redis  redisclient.RedisAdapter
}
