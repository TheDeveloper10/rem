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
		{ "/basic-test", http.MethodPut, http.StatusOK },
		{ "/basic-test", http.MethodPatch, http.StatusOK },
		{ "/basic-test", http.MethodDelete, http.StatusMethodNotAllowed },

		{ "/unknown-path", http.MethodPatch, http.StatusNotFound },
		{ "/basic-test/1", http.MethodGet, http.StatusNotFound },
		{ "/basic-test/2", http.MethodPost, http.StatusNotFound },
		{ "/basic-test/3", http.MethodDelete, http.StatusNotFound },
	}

	runNoBodyRouterTests(t, &tests, router)
}

func createRouter1() *Router {
	router := NewRouter()

	router.
		NewRoute("/%ь")

	router.
		NewRoute("/basic-test").
		Get(func(res IResponse, req IRequest) bool {
			res.Status(http.StatusOK)
			return true
		}).
		Post(func(res IResponse, req IRequest) bool {
			res.Status(http.StatusCreated)
			return true
		}).
		Put(func(res IResponse, req IRequest) bool {
			res.Status(http.StatusOK)
			return true
		}).
		Patch(func(res IResponse, req IRequest) bool {
			res.Status(http.StatusOK)
			return false
		})

	return router
}