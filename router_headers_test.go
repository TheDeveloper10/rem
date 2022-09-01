package rem

import (
	"net/http"
	"testing"
)

func TestRouterHeaders(t *testing.T) {
	router := createRouter4()

	tests := []expectedHeadersRouterTest{
		{ "/headers-test-1", http.MethodGet, http.StatusMethodNotAllowed, nil },
		{ "/headers-test-1/", http.MethodPost, http.StatusOK, map[string]string { "Content-Type": "application/json" } },
	}

	runExpectedHeadersRouterTests(t, &tests, router)
}

func createRouter4() *Router {
	router := NewRouter()

	router.
		NewRoute("/headers-test-1").
		PostRoute(func(res IResponse, req IRequest) bool {
			res.
				Status(http.StatusOK).
				JSON(dummyTestData{ val: 17, str: "testing" })
			return true
		})

	return router
}

type dummyTestData struct {
	val int
	str string
}