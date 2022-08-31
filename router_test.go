package rem

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type routerTest struct {
	id  int
	url string
	method string
	body io.ReadCloser
	expectedStatus int
}

func TestRouter(t *testing.T) {
	router := CreateDefaultRouter()

	router.
	NewBasicRoute("/basic-test").
	GetRoute(func(res IResponse, req IRequest) bool {
		res.Status(http.StatusOK)
		return true
	}).
	PostRoute(func(res IResponse, req IRequest) bool {
		res.Status(http.StatusCreated)
		return true
	})

	tests := []routerTest{
		{ 0, "/basic-test", http.MethodGet, nil, http.StatusOK },
		{ 1, "/basic-test", http.MethodPost, nil, http.StatusCreated },
		{ 2, "/basic-test", http.MethodDelete, nil, http.StatusMethodNotAllowed },
		{ 3, "/unknown-path", http.MethodPatch, nil, http.StatusNotFound },
	}

	for _, test := range tests {
		req, err := http.NewRequest(test.method, test.url, test.body)
		if err != nil {
			panic(err)
		}

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		res := rec.Result()

		if res.StatusCode != test.expectedStatus {
			t.Errorf("TestId: %v\tExpected Status: %v\tReceived Status: %v", test.id, test.expectedStatus, res.StatusCode)
		}
	}
}