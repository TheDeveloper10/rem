package rem

import (
	"net/http"
	"testing"
)

func TestRouterQueryParams(t *testing.T) {
	router := createRouter2()

	tests := []routerTest{
		{ "/query-single-int-param-test", http.MethodGet, nil, http.StatusBadRequest },
		{ "/query-single-int-param-test?v", http.MethodGet, nil, http.StatusBadRequest },
		{ "/query-single-int-param-test?v=", http.MethodGet, nil, http.StatusBadRequest },
		{ "/query-single-int-param-test?v=1", http.MethodGet, nil, http.StatusOK },
		{ "/query-single-int-param-test?v=2", http.MethodGet, nil, http.StatusAccepted },
		{ "/query-single-int-param-test?v=3", http.MethodGet, nil, http.StatusForbidden },
		{ "/query-multi-int-param-test", http.MethodPost, nil, http.StatusBadRequest },
		{ "/query-multi-int-param-test?v1=88", http.MethodPost, nil, http.StatusCreated },
		{ "/query-multi-int-param-test?v2=44", http.MethodPost, nil, http.StatusAccepted },
		{ "/query-multi-int-param-test?v1=15&v2=35", http.MethodPost, nil, http.StatusOK },
	}

	runTests(t, &tests, router)
}

func createRouter2() *Router {
	router := NewRouter()

	router.
		NewRoute("/query-single-int-param-test").
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

	router.
		NewRoute("/query-multi-int-param-test").
		PostRoute(func(res IResponse, req IRequest) bool {
			v1 := req.GetQueryParameters().Get("v1")
			v2 := req.GetQueryParameters().Get("v2")

			if v1 == "15" && v2 == "35" {
				res.Status(http.StatusOK)
			} else if v1 == "88" && v2 == "" {
				res.Status(http.StatusCreated)
			} else if v1 == "" && v2 == "44" {
				res.Status(http.StatusAccepted)
			} else {
				res.Status(http.StatusBadRequest)
			}

			return true
		})

	return router
}