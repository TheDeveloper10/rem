package rem

import (
	"bytes"
	"encoding/json"
	"reflect"
)

type MockResponse struct {
	StatusCode int
	Headers    KeyValues
	Body       []byte
}

func (mr *MockResponse) Status(statusCode int) IResponse {
	mr.StatusCode = statusCode
	return mr
}

func (mr *MockResponse) Header(key string, value string) IResponse {
	mr.Headers.set(key, []string{ value })
	return mr
}

func (mr *MockResponse) Bytes(data []byte) IResponse {
	mr.Body = data
	return mr
}

func (mr *MockResponse) Text(text string) IResponse {
	return mr.Bytes([]byte(text))
}

func (mr *MockResponse) JSON(data interface{}) IResponse {
	bytesData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return mr.Header("Content-Type", "application/json").Bytes(bytesData)
}



func (mr *MockResponse) CompareStatus(other *MockResponse) bool {
	return mr.StatusCode == other.StatusCode
}

func (mr *MockResponse) CompareHeaders(other *MockResponse) bool {
	if len(mr.Headers) != len(other.Headers) {
		return false
	}

	for k, av := range mr.Headers {
		if bv, ok := other.Headers[k]; !ok || !reflect.DeepEqual(av, bv) {
			return false
		}
	}

	return true
}

func (mr *MockResponse) CompareBody(other *MockResponse) bool {
	return bytes.Equal(mr.Body, other.Body)
}
