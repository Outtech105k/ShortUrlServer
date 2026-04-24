package redisclient_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetURLRecord(t *testing.T) {
	id := "id11"
	baseURL := "https://example.com/target"

	mr, adapter, cleanup := setupTestEnvironment(t)

	err := adapter.SetURLRecord(id, baseURL, true, nil)
	assert.NoError(t, err)
	assert.Equal(t, baseURL, mr.HGet(id, "base_url"))

	isSandCushion, err := strconv.ParseBool(mr.HGet(id, "cushion"))
	assert.NoError(t, err)
	assert.Equal(t, true, isSandCushion)

	cleanup()
}
