package utils

import redisclient "github.com/Outtech105k/ShortUrlServer/app-redirect/redis-client"

type AppContext struct {
	Config Config
	Redis  redisclient.RedisAdapter
}
