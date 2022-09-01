package rem

import (
	"net/http"
	"testing"
)

func TestRouterURLParams(t *testing.T) {
	router := createRouter3()

	tests := []noBodyRouterTest{
		{ "/url-test-1", http.MethodGet, http.StatusNotFound },
		{ "/url-test-1/", http.MethodGet, http.StatusNotFound },
		{ "/url-test-1/999", http.MethodGet, http.StatusAccepted },
		{ "/url-test-1/999/", http.MethodGet, http.StatusAccepted },
		{ "/url-test-1/7123", http.MethodGet, http.StatusOK },
		{ "/url-test-1/7123/", http.MethodGet, http.StatusOK },
		{ "/url-test-1/test", http.MethodGet, http.StatusForbidden },
		{ "/url-test-1/test/", http.MethodGet, http.StatusForbidden },

		{ "/url-test-2/", http.MethodGet, http.StatusNotFound },
		{ "/url-test-2/", http.MethodPost, http.StatusNotFound },
		{ "/url-test-2/a/", http.MethodPost, http.StatusNotFound },
		{ "/url-test-2/a/b", http.MethodPost, http.StatusUnprocessableEntity },
		{ "/url-test-2/_/b", http.MethodPost, http.StatusUnprocessableEntity },
		{ "/url-test-2//b", http.MethodPost, http.StatusNotFound },
		{ "/url-test-2//b/", http.MethodPost, http.StatusNotFound },
		{ "/url-test-2/qwe/b", http.MethodPost, http.StatusUnprocessableEntity },
		{ "/url-test-2/qwe/b/", http.MethodPost, http.StatusUnprocessableEntity },
		{ "/url-test-2/qwe/ewq", http.MethodPost, http.StatusOK },
		{ "/url-test-2/qwe/ewq", http.MethodGet, http.StatusMethodNotAllowed },
		{ "/url-test-2/qwe/ewq/", http.MethodPost, http.StatusOK },
		{ "/url-test-2/1/2", http.MethodPost, http.StatusAccepted },
		{ "/url-test-2/1/2/", http.MethodPost, http.StatusAccepted },
		{ "/url-test-2/2/1", http.MethodPost, http.StatusUnprocessableEntity },
		{ "/url-test-2/2/1/", http.MethodPost, http.StatusUnprocessableEntity },

		{ "/url-test-3/", http.MethodGet, http.StatusNotFound },
		{ "/url-test-3/", http.MethodDelete, http.StatusNotFound },
		{ "/url-test-3//data/", http.MethodDelete, http.StatusNotFound },
		{ "/url-test-3/data/data", http.MethodDelete, http.StatusUnprocessableEntity },
		{ "/url-test-3/data/data", http.MethodGet, http.StatusMethodNotAllowed },
		{ "/url-test-3/data/data2", http.MethodDelete, http.StatusNotFound },
		{ "/url-test-3/data/adata", http.MethodDelete, http.StatusNotFound },
		{ "/url-test-3/123/data", http.MethodDelete, http.StatusOK },
		{ "/url-test-3/123//data///", http.MethodDelete, http.StatusOK },
		{ "/url-test-3/h///data", http.MethodDelete, http.StatusUnauthorized },
		{ "/url-test-3//h/data", http.MethodDelete, http.StatusUnauthorized },
		{ "/url-test-3///h/data", http.MethodDelete, http.StatusUnauthorized },
		{ "/url-test-3/test12345test/data", http.MethodDelete, http.StatusAccepted },
		{ "/url-test-3/test12345test/data/", http.MethodDelete, http.StatusAccepted },
	}

	// test for route /a/b/:

	runNoBodyRouterTests(t, &tests, router)
}

func createRouter3() *Router {
	router := NewRouter()

	router.
		NewRoute("/url-test-1/:testId").
		GetRoute(func(res IResponse, req IRequest) bool {
			testId, status := req.GetURLParameters().Get("testId")
			if !status {
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
		NewRoute("/url-test-2/:a/:g/").
		PostRoute(func(res IResponse, req IRequest) bool {
			a, statusA := req.GetURLParameters().Get("a")
			if !statusA {
				res.Status(http.StatusBadRequest)
				return true
			}
			g, statusG := req.GetURLParameters().Get("g")
			if !statusG {
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
		NewRoute("/url-test-3/:userId/data/").
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