package utils

import (
	redisclient "github.com/Outtech105k/ShortUrlServer/app-ctl/redis-client"
)

type AppCopntext struct {
	Redis redisclient.RedisAdapter
}
