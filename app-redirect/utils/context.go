package utils

import redisclient "github.com/Outtech105k/ShortUrlServer/app-redirect/redis-client"

type AppContext struct {
	Redis redisclient.RedisAdapter
}
