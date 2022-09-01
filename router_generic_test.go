package rem

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type routerTest interface {
	PerformTest(int, *Router) error
}

// -----------------------------------------------------------
// No Body request Router Test
// -----------------------------------------------------------

// Router Test that has no body in the request
type noBodyRouterTest struct {
	url 		   string
	method 		   string
	expectedStatus int
}

func (nbrt *noBodyRouterTest) PerformTest(testId int, router *Router) error {
	req, err := http.NewRequest(nbrt.method, nbrt.url, nil)
	if err != nil {
		return err
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	res := rec.Result()

	if res.StatusCode != nbrt.expectedStatus {
		return fmt.Errorf("TestId: %v\tURL: %v\tExpected Status: %v\tReceived Status: %v", testId, nbrt.url, nbrt.expectedStatus, res.StatusCode)
	}

	return nil
}

func runNoBodyRouterTests(t *testing.T, tests *[]noBodyRouterTest, router *Router) {
	for testId, test := range *tests {
		err := test.PerformTest(testId, router)
		if err != nil {
			t.Error(err)
		}
	}
}


// -----------------------------------------------------------
// Expected Headers response Router Test
// -----------------------------------------------------------

type expectedHeadersRouterTest struct {
	url 			string
	method 			string
	expectedStatus int
	expectedHeaders map[string]string
}

func (ehrt *expectedHeadersRouterTest) PerformTest(testId int, router *Router) error {
	req, err := http.NewRequest(ehrt.method, ehrt.url, nil)
	if err != nil {
		return err
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	res := rec.Result()

	if res.StatusCode != ehrt.expectedStatus {
		return fmt.Errorf("TestId: %v\tURL: %v\tExpected Status: %v\tReceived Status: %v", testId, ehrt.url, ehrt.expectedStatus, res.StatusCode)
	}

	for k, v := range ehrt.expectedHeaders {
		header := res.Header.Get(k)
		if header == "" {
			return fmt.Errorf("TestId: %v\tExpected Header: '%v: %v'\tReceived an empty header", testId, k, v)
		} else if header != v {
			return fmt.Errorf("TestId: %v\tExpected Header: '%v: %v'\tReceived Header: '%v: %v'", testId, k, v, k, header)
		}
	}

	return nil
}

func runExpectedHeadersRouterTests(t *testing.T, tests *[]expectedHeadersRouterTest, router *Router) {
	for testId, test := range *tests {
		err := test.PerformTest(testId, router)
		if err != nil {
			t.Error(err)
		}
	}
}
