package controllers_test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/Outtech105k/ShortUrlServer/app/models"
)

func TestSetUrlHandler_Integration(t *testing.T) {
	appCtx, mr, router, cleanup := setupTestEnvironment(t)
	defer cleanup()

	t.Run("Create Random URL", func(t *testing.T) {
		reqBody := models.SetUrlRequest{
			BaseURL: "https://example.com/original",
		}
		w := performRequest(router, "POST", "/set", reqBody)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		var resp models.APIResponce
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Errorf("responce json unmarshal: %v", err)
		}

		short := strings.TrimPrefix(resp.ShortURL, appCtx.Config.ServerEndpoint+"/")
		mrGot := mr.HGet(short, "base_url")
		if mrGot != resp.BaseURL {
			t.Errorf("expected base_url %s, got %s", resp.BaseURL, mrGot)
		}
	})

	t.Run("Success with Custom ID", func(t *testing.T) {
		customId := "my-id"
		reqBody := models.SetUrlRequest{
			BaseURL:  "https://example.com/custom",
			CustomID: &customId,
		}
		w := performRequest(router, "POST", "/set", reqBody)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		if !mr.Exists(customId) {
			t.Errorf("expected key %s to exist in redis", customId)
		}
		mrGot := mr.HGet(customId, "base_url")
		if mrGot != reqBody.BaseURL {
			t.Errorf("expected base_url %s, got %s", reqBody.BaseURL, mrGot)
		}
	})

	t.Run("Conflict with Existing ID", func(t *testing.T) {
		existingId := "already-used"
		mr.HSet(existingId, "base_url", "https://existing.com")

		reqBody := models.SetUrlRequest{
			BaseURL:  "https://example.com/new",
			CustomID: &existingId,
		}
		w := performRequest(router, "POST", "/set", reqBody)

		if w.Code != http.StatusConflict {
			t.Errorf("expected status 409, got %d", w.Code)
		}
	})

	t.Run("Validation Error (Invalid URL)", func(t *testing.T) {
		reqBody := models.SetUrlRequest{
			BaseURL: "not-a-url",
		}
		w := performRequest(router, "POST", "/set", reqBody)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", w.Code)
		}
	})
}
