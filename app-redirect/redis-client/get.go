package redisclient

func (r *RedisAdapter) GetBaseUrl(key string) (string, error) {
	baseUrl, err := r.client.HGet(key, "base_url").Result()
	if err != nil {
		return "", err
	}

	return baseUrl, nil
}
