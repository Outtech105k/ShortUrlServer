package redisclient_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBaseUrl(t *testing.T) {
	key := "id1"
	baseUrl := "https://example.com/target"
	mr, adapter, cleanup := setupTestEnvironment(t)

	mr.HSet(
		key,
		"base_url",
		baseUrl,
		"cushion",
		"false",
	)

	result, err := adapter.GetBaseUrl(key)
	assert.NoError(t, err)
	assert.Equal(t, baseUrl, result)

	cleanup()
}

func TestGetIsNeedCusionPage(t *testing.T) {
	key := "id2"
	baseUrl := "https://example.com/target"
	mr, adapter, cleanup := setupTestEnvironment(t)

	mr.HSet(
		key,
		"base_url",
		baseUrl,
		"cushion",
		"true",
	)

	result, err := adapter.GetIsNeedCusionPage(key)
	assert.NoError(t, err)
	assert.Equal(t, true, result)
	cleanup()
}
