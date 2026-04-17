package redisclient

import "fmt"

func (r *RedisAdapter) GetBaseUrl(key string) (string, error) {
	baseUrl, err := r.client.HGet(key, "base_url").Result()
	if err != nil {
		return "", err
	}

	return baseUrl, nil
}

func (r *RedisAdapter) GetIsNeedCusionPage(key string) (bool, error) {
	isNeedCusionPage, err := r.client.HGet(key, "cushion").Result()
	if err != nil {
		return false, err
	}

	if isNeedCusionPage == "1" {
		return true, nil
	} else if isNeedCusionPage == "0" {
		return false, nil
	}

	return false, fmt.Errorf("invalid value for cushion: %s", isNeedCusionPage)
}
