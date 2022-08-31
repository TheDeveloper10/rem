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
	expectedStatus int
}

func runTests(t *testing.T, tests *[]routerTest, router *Router) {
	for testId, test := range *tests {
		req, err := http.NewRequest(test.method, test.url, test.body)
		if err != nil {
			panic(err)
		}

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		res := rec.Result()

		if res.StatusCode != test.expectedStatus {
			t.Errorf("TestId: %v\tURL: %v\tExpected Status: %v\tReceived Status: %v", testId, test.url, test.expectedStatus, res.StatusCode)
		}
	}
}

func TestRouterBasic(t *testing.T) {
	router := createRouter1()

	tests := []routerTest{
		{ "/basic-test", http.MethodGet, nil, http.StatusOK },
		{ "/basic-test", http.MethodPost, nil, http.StatusCreated },
		{ "/basic-test", http.MethodDelete, nil, http.StatusMethodNotAllowed },
		{ "/unknown-path", http.MethodPatch, nil, http.StatusNotFound },
	}

	runTests(t, &tests, router)
}

func createRouter1() *Router {
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

	return router
}