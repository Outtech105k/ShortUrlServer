package redisclient_test

import (
	"testing"

	redisclient "github.com/Outtech105k/ShortUrlServer/app/redis-client"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
)

func setupTestEnvironment(t *testing.T) (*miniredis.Miniredis, redisclient.RedisAdapter, func()) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("failed to run miniredis: %v", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	adapter := &redisclient.RedisAdapter{
		Client: client,
	}

	cleanup := func() {
		adapter.Close()
		mr.Close()
	}

	return mr, *adapter, cleanup
}
