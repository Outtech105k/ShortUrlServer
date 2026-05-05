package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	redisclient "github.com/Outtech105k/ShortUrlServer/app/redis-client"
	"github.com/Outtech105k/ShortUrlServer/app/routes"
	"github.com/Outtech105k/ShortUrlServer/app/utils"
	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// MockRedisClient の定義をここに追加
type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) SetURLRecord(id string, baseUrl string, isSandCushion bool, expireDelta *time.Duration) error {
	args := m.Called(id, baseUrl, isSandCushion, expireDelta)
	return args.Error(0)
}

func (m *MockRedisClient) GetBaseUrl(key string) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

func (m *MockRedisClient) GetIsNeedCusionPage(key string) (bool, error) {
	args := m.Called(key)
	return args.Bool(0), args.Error(1)
}

func (m *MockRedisClient) IsExists(key string) (bool, error) {
	args := m.Called(key)
	return args.Bool(0), args.Error(1)
}

func (m *MockRedisClient) Close() error {
	args := m.Called()
	return args.Error(0)
}

// テスト用の共通環境を構築
// 戻り値: コンテキスト, miniredisインスタンス, ルーター, クリーンアップ関数
func setupTestEnvironment(t *testing.T) (utils.AppContext, *miniredis.Miniredis, *gin.Engine, func()) {
	t.Helper()

	// templatesディレクトリへのパス調整 (app/controllers から app/ へ移動)
	oldWd, _ := os.Getwd()
	if err := os.Chdir(".."); err != nil {
		t.Fatalf("failed to change directory: %v", err)
	}

	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("failed to run miniredis: %v", err)
	}

	adapter, err := redisclient.NewRedisAdapter(mr.Addr())
	if err != nil {
		t.Fatalf("failed to create redis adapter: %v", err)
	}

	appCtx := &utils.AppContext{
		Config: utils.Config{
			ServerEndpoint: "https://srv.test",
		},
		Redis: adapter,
	}

	router := routes.SetupRouter(appCtx)

	cleanup := func() {
		adapter.Close()
		mr.Close()
		if err := os.Chdir(oldWd); err != nil {
			t.Logf("Warning: failed to restore directory: %v", err)
		}
	}

	return *appCtx, mr, router, cleanup
}

// HTTPリクエストを実行してレスポンスを返す
func performRequest(router *gin.Engine, method, path string, body interface{}) *httptest.ResponseRecorder {
	var buf *bytes.Buffer
	if body != nil {
		b, _ := json.Marshal(body)
		buf = bytes.NewBuffer(b)
	} else {
		buf = bytes.NewBuffer(nil)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, buf)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	router.ServeHTTP(w, req)
	return w
}
