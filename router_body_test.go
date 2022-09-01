package rem

import (
	"net/http"
	"testing"
)

func TestRouterBody(t *testing.T) {
	router := createRouter5()

	tests := []bodyRouterTest{
		{ "/body-test-1/", http.MethodDelete, http.StatusNotFound, "" },
		{
			"/body-test-1/tst",
			http.MethodDelete,
			http.StatusOK,
			"{\"userId\":\"atst\",\"age\":93,\"username\":\"cool_username\",\"lastLoginTime\":17777}",
		},
		{
			"/body-test-1/q",
			http.MethodDelete,
			http.StatusOK,
			"{\"userId\":\"aq\",\"age\":91,\"username\":\"cool_username\",\"lastLoginTime\":17777}",
		},
		{
			"body-test-1/abcd",
			http.MethodDelete,
			http.StatusForbidden,
			"{\"uid\":\"abcdb\",\"test2\":[\"123\",\"abcdc\",\"hello\"]}",
		},
		{
			"body-test-1/abcd",
			http.MethodPost,
			http.StatusForbidden,
			"{\"uid\":\"abcdb\",\"test2\":[\"123\",\"abcdc\",\"hello\"]}",
		},
		{
			"body-test-1/abcd",
			http.MethodPut,
			http.StatusMethodNotAllowed,
			"",
		},
	}

	runBodyRouterTests(t, &tests, router)
}

func createRouter5() *Router {
	router := NewRouter()

	router.
		NewRoute("/body-test-1/:userId").
		MultiMethod([]string{http.MethodDelete, http.MethodPost}, func(res IResponse, req IRequest) bool {
			userId := req.GetURLParameters().Get("userId")
			if userId == "" {
				res.Status(http.StatusBadRequest)
				return true
			}

			if len(userId) <= 3 {
				res.
					Status(http.StatusOK).
					JSON(dummyBodyTestData{
						UserId: "a" + userId,
						Age: 90 + len(userId),
						Username: "cool_username",
						LastLoginTime: 17777,
					})
			} else {
				res.
					JSON(dummyBodyTestData2{
						UID: userId + "b",
						Test: []string{
							"123",
							userId + "c",
							"hello",
						},
					}).
					Status(http.StatusForbidden)
			}

			return true
		})

	return router
}

type dummyBodyTestData struct {
	UserId 	      string `json:"userId"`
	Age	 	      int	 `json:"age"`
	Username      string `json:"username"`
	LastLoginTime int 	 `json:"lastLoginTime"`
}

type dummyBodyTestData2 struct {
	UID  string   `json:"uid"`
	Test []string `json:"test2"`
}