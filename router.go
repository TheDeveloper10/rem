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

		// if the response has a body
		if rBody := response.GetWrittenBody(); rBody != nil {
			// serialize the body
			data, err := Serialize(response.GetWrittenBody())
			if err != nil {
				panic(err)
			}

			// write it to ResponseWriter
			_, err = rw.Write(data)
			if err != nil {
				panic(err)
			}
		}

		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}
