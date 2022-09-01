package rem

import (
	"net/http"
	"testing"
)

func TestRouterBasic(t *testing.T) {
	router := createRouter1()

	tests := []noBodyRouterTest{
		{ "/basic-test", http.MethodGet, http.StatusOK },
		{ "/basic-test", http.MethodPost, http.StatusCreated },
		{ "/basic-test", http.MethodDelete, http.StatusMethodNotAllowed },
		{ "/unknown-path", http.MethodPatch, http.StatusNotFound },
	}

	runNoBodyRouterTests(t, &tests, router)
}

func createRouter1() *Router {
	router := NewRouter()

	router.
		NewRoute("/basic-test").
		Get(func(res IResponse, req IRequest) bool {
			res.Status(http.StatusOK)
			return true
		}).
		Post(func(res IResponse, req IRequest) bool {
			res.Status(http.StatusCreated)
			return true
		})

	return router
}