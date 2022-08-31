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
	router := createRouter()

	tests := []routerTest{
		{ 0, "/basic-test", http.MethodGet, nil, http.StatusOK },
		{ 1, "/basic-test", http.MethodPost, nil, http.StatusCreated },
		{ 2, "/basic-test", http.MethodDelete, nil, http.StatusMethodNotAllowed },
		{ 3, "/unknown-path", http.MethodPatch, nil, http.StatusNotFound },
		{ 4, "/query-single-int-param-test", http.MethodGet, nil, http.StatusBadRequest },
		{ 5, "/query-single-int-param-test?v", http.MethodGet, nil, http.StatusBadRequest },
		{ 6, "/query-single-int-param-test?v=", http.MethodGet, nil, http.StatusBadRequest },
		{ 7, "/query-single-int-param-test?v=1", http.MethodGet, nil, http.StatusOK },
		{ 8, "/query-single-int-param-test?v=2", http.MethodGet, nil, http.StatusAccepted },
		{ 9, "/query-single-int-param-test?v=3", http.MethodGet, nil, http.StatusForbidden },
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

func createRouter() *Router {
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

	router.
		NewBasicRoute("/query-single-int-param-test").
		GetRoute(func(res IResponse, req IRequest) bool {
			v := req.GetQueryParameters().Get("v")

			if v == "" {
				res.Status(http.StatusBadRequest)
			} else if v == "1" {
				res.Status(http.StatusOK)
			} else if v == "2" {
				res.Status(http.StatusAccepted)
			} else {
				res.Status(http.StatusForbidden)
			}

			return true
		})

	//router.
	//	NewBasicRoute("/query-multi-int-param-test").
	//	PostRoute(func(res IResponse, req IRequest) bool {
	//
	//		return true
	//	})

	return router
}