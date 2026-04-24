package utils

type Config struct {
	ServerEndpoint string `env:"ENDPOINT"`
	RedisAddr      string `env:"REDIS_ADDR" envDefault:"redis:6379"`
}
