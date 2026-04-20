package redisclient

import (
	"fmt"
	"strconv"
)

func (r *RedisAdapter) GetBaseUrl(key string) (string, error) {
	baseUrl, err := r.client.HGet(key, "base_url").Result()
	if err != nil {
		return "", err
	}

	return baseUrl, nil
}

func (r *RedisAdapter) GetIsNeedCusionPage(key string) (bool, error) {
	redisVal, err := r.client.HGet(key, "cushion").Result()
	if err != nil {
		return false, fmt.Errorf("get redis: %w", err)
	}

	isNeed, err := strconv.ParseBool(redisVal)
	if err != nil {
		return false, fmt.Errorf("parse got val: %w", err)
	}

	return isNeed, nil
}
