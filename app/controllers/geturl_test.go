package controllers_test

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetUrlHandler_Integration(t *testing.T) {
	_, mr, router, cleanup := setupTestEnvironment(t)
	defer cleanup()

	t.Run("Redirect Without Cushion", func(t *testing.T) {
		shortUrl := "no-cushion"
		baseUrl := "https://example.com/target"

		mr.HSet(shortUrl, "base_url", baseUrl)
		mr.HSet(shortUrl, "cushion", "false")

		w := performRequest(router, "GET", "/"+shortUrl, nil)

		if w.Code != http.StatusFound {
			t.Errorf("expected status 302, got %d", w.Code)
		}
		if w.Header().Get("Location") != baseUrl {
			t.Errorf("expected redirect to %s, got %s", baseUrl, w.Header().Get("Location"))
		}
	})

	t.Run("Show Cushion Page", func(t *testing.T) {
		shortUrl := "with-cushion"
		baseUrl := "https://example.com/cushion-target"

		mr.HSet(shortUrl, "base_url", baseUrl)
		mr.HSet(shortUrl, "cushion", "true")

		w := performRequest(router, "GET", "/"+shortUrl, nil)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}
		if !strings.Contains(w.Body.String(), baseUrl) {
			t.Errorf("expected body to contain %s", baseUrl)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		w := performRequest(router, "GET", "/unknown", nil)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", w.Code)
		}
	})
}
