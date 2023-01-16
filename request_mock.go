package rem

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type MockRequest struct {
	URL string
	Method string
	Headers KeyValues
	Cookies []*http.Cookie
	URLParameters KeyValue
	QueryParameters KeyValues

	// Fill in at most one of them
	// Leave the rest empty
	Body string // body that is a string
	BodyJSON interface{} // object that will be converted to JSON and used as a body
}

func (mr *MockRequest) GetURL() string { return mr.URL }
func (mr *MockRequest) GetMethod() string { return mr.Method }
func (mr *MockRequest) GetHeaders() KeyValues { return mr.Headers }
func (mr *MockRequest) GetCookies() []*http.Cookie { return mr.Cookies }
func (mr *MockRequest) GetURLParameters() KeyValue { return mr.URLParameters }
func (mr *MockRequest) setURLParameters(params *KeyValue) { mr.URLParameters = *params }
func (mr *MockRequest) GetQueryParameters() KeyValues { return mr.QueryParameters }
func (mr *MockRequest) GetBody() io.ReadCloser {
	body := ""
	if mr.BodyJSON != nil {
		bodyBytes, err := json.Marshal(mr.BodyJSON)
		if err != nil {
			panic(err)
		}
		body = string(bodyBytes)
	} else {
		body = mr.Body
	}

	readCloser := ioutil.NopCloser(strings.NewReader(body))
	return readCloser
}
