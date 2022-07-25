package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Bakhram74/rest-api.git/internal/app/model"
	"github.com/Bakhram74/rest-api.git/internal/app/store/teststore"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_AuthenticationUser(t *testing.T) {
	store := teststore.New()
	u := model.TestingUser(t)
	store.User().Create(u)
	testcase := []struct {
		name        string
		cookieValue map[interface{}]interface{}
		expected    int
	}{
		{
			name: "authenticated",
			cookieValue: map[interface{}]interface{}{
				"user_id": u.ID,
			},
			expected: http.StatusOK,
		},
		{
			name:        "unauthenticated",
			cookieValue: nil,
			expected:    http.StatusUnauthorized,
		},
	}
	secretKey := []byte("secret")
	s := NewServer(store, sessions.NewCookieStore(secretKey))
	sc := securecookie.New(secretKey, nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			cookieStr, _ := sc.Encode(sessionName, tc.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
			s.authenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, tc.expected, rec.Code)
		})
	}
}
func TestServer_HandleUsersCreate(t *testing.T) {
	s := NewServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	testCase := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "example@mail.com",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		}, {
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			r, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, r)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestingUser(t)
	store := teststore.New()
	store.User().Create(u)
	s := NewServer(store, sessions.NewCookieStore([]byte("secret")))
	testCase := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:    "invalid",
			payload: "invalid",

			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})

	}
}
