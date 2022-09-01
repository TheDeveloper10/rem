package rem

import (
	"net/http"
	"testing"
)

func TestRouterHeaders(t *testing.T) {
	router := createRouter4()

	tests := []headersRouterTest{
		{ "/headers-test-1", http.MethodGet, http.StatusMethodNotAllowed, nil },
		{ "/headers-test-1/", http.MethodPost, http.StatusOK, map[string]string { "Content-Type": "application/json" } },
		{ "/headers-test-1/17j/", http.MethodPost, http.StatusCreated, map[string]string { "Content-Type": "application/json", "Content-Encoding": "gzip" } },
		{ "/headers-test-1/z", http.MethodPost, http.StatusUnauthorized, map[string]string { "Access-Control-Allow-Origin": "*", "Last-Modified": "yesterday" } },
		{ "/headers-test-1/f", http.MethodPost, http.StatusForbidden, map[string]string{} },
	}

	runHeadersRouterTest(t, &tests, router)
}

func createRouter4() *Router {
	router := NewRouter()

	router.
		NewRoute("/headers-test-1").
		PostRoute(func(res IResponse, req IRequest) bool {
			res.
				Status(http.StatusOK).
				JSON(dummyTestData{ Val: 17, Str: "testing" })
			return true
		})

	router.
		NewRoute("/headers-test-1/:userId/").
		PostRoute(func(res IResponse, req IRequest) bool {
			userId, status := req.GetURLParameters().Get("userId")
			if !status {
				res.Status(http.StatusBadRequest)
				return true
			}

			if userId == "17j" {
				res.
					Status(http.StatusCreated).
					Header("Content-Encoding", "gzip").
					JSON(dummyTestData{Str: "123", Val: 12})
			} else if userId == "z" {
				res.
					Header("Access-Control-Allow-Origin", "*").
					Text("hi!").
					Status(http.StatusUnauthorized).
					Header("Last-Modified", "yesterday") // I know it's not according to the standard. I just want it to be shorter
			} else {
				res.Status(http.StatusForbidden)
			}

			return true
		})

	return router
}

type dummyTestData struct {
	Val int    `json:"val"`
	Str string `json:"str"`
}