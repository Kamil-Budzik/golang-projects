package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kamil-budzik/gin-playground/internal/routes"
)

func TestPing(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}
}
