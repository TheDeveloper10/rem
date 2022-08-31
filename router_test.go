package rem

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type routerTest struct {
	url string
	method string
	body io.ReadCloser
}

func TestRouter(t *testing.T) {
	tests := []routerTest{
		{ "/test", http.MethodGet, nil },
	}

	router := CreateDefaultRouter()
	router.GetRoute("/test", func(res IResponse, req IRequest) bool {
		res.Status(http.StatusOK)
		return true
	})

	for _, test := range tests {
		req, err := http.NewRequest(test.method, test.url, test.body)
		if err != nil {
			panic(err)
		}

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		res := rec.Result()
		t.Errorf("Status: %v", res.StatusCode)
	}
}