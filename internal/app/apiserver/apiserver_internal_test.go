package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	config := Config{
		BindAddr: os.Getenv("bind_addr"),
		LogLevel: os.Getenv("log_level"),
	}
	s := New(&config)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, "hello", rec.Body.String())
}
