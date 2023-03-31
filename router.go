package rem

import "net/http"

type Router struct {
	routes []IRoute
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	targetURL := cleanPath(req.URL.EscapedPath())

	for _, route := range r.routes {
		if !route.Match(targetURL) {
			continue
		}

		request := wrapHTTPRequest(req)
		ctx := newContext(request, route.GetHandlers())
		response := ctx.Next()

		// write status
		rw.WriteHeader(response.GetWrittenStatus())

		// set headers (if any exist)
		for header, value := range response.GetWrittenHeaders() {
			rw.Header().Set(header, value)
		}

		// if the response doesn't have a body
		if rBody := response.GetWrittenBody(); rBody == nil {
			return
		}

		// serialize the body
		data, err := config.BodyProcessor.SerializeResponse(response)
		if err != nil {
			panic(err)
		}

		// write it to ResponseWriter
		_, err = rw.Write(data)
		if err != nil {
			panic(err)
		}

		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}
