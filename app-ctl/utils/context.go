package utils

import (
	redisclient "github.com/Outtech105k/ShortUrlServer/app-ctl/redis-client"
)

type AppContext struct {
	Redis redisclient.RedisAdapter
}
