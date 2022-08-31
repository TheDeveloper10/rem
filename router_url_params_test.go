package rem

import (
	"net/http"
	"testing"
)

func TestRouterURLParams(t *testing.T) {
	router := createRouter3()

	tests := []routerTest{
		{ "/url-test-1", http.MethodGet, nil, http.StatusNotFound },
		{ "/url-test-1/", http.MethodGet, nil, http.StatusNotFound },
		{ "/url-test-1/999", http.MethodGet, nil, http.StatusAccepted },
		{ "/url-test-1/999/", http.MethodGet, nil, http.StatusAccepted },
		{ "/url-test-1/7123", http.MethodGet, nil, http.StatusOK },
		{ "/url-test-1/7123/", http.MethodGet, nil, http.StatusOK },
		{ "/url-test-1/test", http.MethodGet, nil, http.StatusForbidden },
		{ "/url-test-1/test/", http.MethodGet, nil, http.StatusForbidden },

		{ "/url-test-2/", http.MethodGet, nil, http.StatusNotFound },
		{ "/url-test-2/", http.MethodPost, nil, http.StatusNotFound },
		{ "/url-test-2/a/", http.MethodPost, nil, http.StatusNotFound },
		{ "/url-test-2/a/b", http.MethodPost, nil, http.StatusUnprocessableEntity },
		{ "/url-test-2/_/b", http.MethodPost, nil, http.StatusUnprocessableEntity },
		{ "/url-test-2//b", http.MethodPost, nil, http.StatusNotFound },
		{ "/url-test-2//b/", http.MethodPost, nil, http.StatusNotFound },
		{ "/url-test-2/qwe/b", http.MethodPost, nil, http.StatusUnprocessableEntity },
		{ "/url-test-2/qwe/b/", http.MethodPost, nil, http.StatusUnprocessableEntity },
		{ "/url-test-2/qwe/ewq", http.MethodPost, nil, http.StatusOK },
		{ "/url-test-2/qwe/ewq", http.MethodGet, nil, http.StatusMethodNotAllowed },
		{ "/url-test-2/qwe/ewq/", http.MethodPost, nil, http.StatusOK },
		{ "/url-test-2/1/2", http.MethodPost, nil, http.StatusAccepted },
		{ "/url-test-2/1/2/", http.MethodPost, nil, http.StatusAccepted },
		{ "/url-test-2/2/1", http.MethodPost, nil, http.StatusUnprocessableEntity },
		{ "/url-test-2/2/1/", http.MethodPost, nil, http.StatusUnprocessableEntity },

		{ "/url-test-3/", http.MethodGet, nil, http.StatusNotFound },
		{ "/url-test-3/", http.MethodDelete, nil, http.StatusNotFound },
		{ "/url-test-3//data/", http.MethodDelete, nil, http.StatusNotFound },
		{ "/url-test-3/data/data", http.MethodDelete, nil, http.StatusUnprocessableEntity },
		{ "/url-test-3/data/data", http.MethodGet, nil, http.StatusMethodNotAllowed },
		{ "/url-test-3/data/data2", http.MethodDelete, nil, http.StatusNotFound },
		{ "/url-test-3/data/adata", http.MethodDelete, nil, http.StatusNotFound },
		{ "/url-test-3/123/data", http.MethodDelete, nil, http.StatusOK },
		{ "/url-test-3/123//data///", http.MethodDelete, nil, http.StatusOK },
		{ "/url-test-3/h///data", http.MethodDelete, nil, http.StatusUnauthorized },
		{ "/url-test-3//h/data", http.MethodDelete, nil, http.StatusUnauthorized },
		{ "/url-test-3///h/data", http.MethodDelete, nil, http.StatusUnauthorized },
		{ "/url-test-3/test12345test/data", http.MethodDelete, nil, http.StatusAccepted },
		{ "/url-test-3/test12345test/data/", http.MethodDelete, nil, http.StatusAccepted },
	}

	// test for route /a/b/:

	runTests(t, &tests, router)
}

func createRouter3() *Router {
	router := CreateDefaultRouter()

	router.
		NewVariableRoute("/url-test-1/:testId").
		GetRoute(func(res IResponse, req IRequest) bool {
			testId, ok := req.GetURLParameters()["testId"]
			if !ok {
				res.Status(http.StatusBadRequest)
				return true
			}

			if testId == "7123" {
				res.Status(http.StatusOK)
			} else if testId == "test" {
				res.Status(http.StatusForbidden)
			} else {
				res.Status(http.StatusAccepted)
			}

			return true
		})

	router.
		NewVariableRoute("/url-test-2/:a/:g/").
		PostRoute(func(res IResponse, req IRequest) bool {
			a, ok := req.GetURLParameters()["a"]
			if !ok {
				res.Status(http.StatusBadRequest)
				return true
			}
			g, ok := req.GetURLParameters()["g"]
			if !ok {
				res.Status(http.StatusBadRequest)
				return true
			}

			if a == "" {
				res.Status(http.StatusUnauthorized)
			} else if g == "" {
				res.Status(http.StatusForbidden)
			} else if a == "qwe" && g == "ewq" {
				res.Status(http.StatusOK)
			} else if a == "1" && g == "2" {
				res.Status(http.StatusAccepted)
			} else {
				res.Status(http.StatusUnprocessableEntity)
			}

			return true
		})

	router.
		NewVariableRoute("/url-test-3/:userId/data/").
		DeleteRoute(func(res IResponse, req IRequest) bool {
			userId, ok := req.GetURLParameters()["userId"]
			if !ok {
				res.Status(http.StatusBadRequest)
				return true
			}

			if userId == "123" {
				res.Status(http.StatusOK)
			} else if userId == "h" {
				res.Status(http.StatusUnauthorized)
			} else if userId == "test12345test" {
				res.Status(http.StatusAccepted)
			} else {
				res.Status(http.StatusUnprocessableEntity)
			}

			return true
		})

	return router
}