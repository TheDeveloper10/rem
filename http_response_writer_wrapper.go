package rem

import (
	"encoding/json"
	"net/http"
)

// A wrapper of http.ResponseWriter that implements IResponse
type HTTPResponseWriterWrapper struct {
	rw http.ResponseWriter
	statusCode int
	body       []byte
}

// Writes a status code
func (w *HTTPResponseWriterWrapper) Status(statusCode int) IResponse {
	w.statusCode = statusCode
	return w
}

// Set a new header
// (Writes directly to http.ResponseWriter because the only thing that
// matters is the headers to be written first)
func (w *HTTPResponseWriterWrapper) Header(key string, value string) IResponse {
	w.rw.Header().Set(key, value)
	return w
}

// Write bytes to the body
func (w *HTTPResponseWriterWrapper) Bytes(bytes []byte) IResponse {
	w.body = bytes
	return w
}

// Write text to the body
func (w *HTTPResponseWriterWrapper) Text(text string) IResponse {
	return w.Bytes([]byte(text))
}

// Write JSON to the body (also sets Content-Type header)
func (w *HTTPResponseWriterWrapper) JSON(data interface{}) IResponse {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return w.Header("Content-Type", "application/json").Bytes(bytes)
}

// This method is required because http.ResponseWriter is obsessed over
// the order in which we give the headers, status code and body.
func (w *HTTPResponseWriterWrapper) flush() {
	w.rw.WriteHeader(w.statusCode)

	if w.body != nil {
		_, err := w.rw.Write(w.body)
		if err != nil {
			panic(err)
		}
	}
}