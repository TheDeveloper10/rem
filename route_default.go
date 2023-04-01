package rem

import (
	"net/http"
	"regexp"
)

func newRouteDefault(url string) IRoute {
	cleanedURL := cleanPath(url)
	endpoints := map[string][]Handler{}

	hasMatched, err := regexp.Match(":\\w+", []byte(url))
	if err != nil {
		panic(err)
	}

	return &routeDefault{
		url:          cleanedURL,
		urlLen:       len(url),
		endpoints:    endpoints,
		hasVariables: hasMatched,
	}
}

type routeDefault struct {
	url       string
	urlLen    int
	endpoints map[string][]Handler
	// whether the route has any variables or not
	hasVariables bool
}

func (rd *routeDefault) SetHandlers(method string, handlers []Handler) IRoute {
	rd.endpoints[method] = handlers
	return rd
}

func (rd *routeDefault) Match(targetURL string) bool {
	if !rd.hasVariables {
		return rd.url == targetURL
	}

	tLen := len(targetURL)

	oIndx := 0
	tIndx := 0
	for {
		if rd.url[oIndx] == ':' {
			tIndx = skipToNextIndent(targetURL, tLen, tIndx)
			oIndx = skipToNextIndent(rd.url, rd.urlLen, oIndx)
		} else if rd.url[oIndx] != targetURL[tIndx] {
			return false
		}

		oIndx++
		tIndx++

		oOverTop := oIndx >= rd.urlLen
		tOverTop := tIndx >= tLen

		if oOverTop && tOverTop {
			return true
		} else if !oOverTop && !tOverTop {
			continue
		} else {
			return false
		}
	}
}

func (rd *routeDefault) ExtractURLParameters(targetURL string) (map[string]string, bool) {
	if rd.hasVariables {
		return nil, true
	}

	var (
		params = map[string]string{}
		tLen   = len(targetURL)
		oIndx  = 0 // index on original URL
		tIndx  = 0 // index on target URL
	)

	for oIndx < rd.urlLen && tIndx < tLen {
		if rd.url[oIndx] == ':' {
			variable := extractToNextIndent(targetURL, tLen, tIndx)
			name := extractToNextIndent(rd.url, rd.urlLen, oIndx+1)
			tIndx = tIndx + len(variable)
			oIndx = oIndx + len(name) + 1
			params[name] = variable
		} else if rd.url[oIndx] != targetURL[tIndx] { // mismatch between original and actual
			return nil, false
		} else {
			oIndx++
			tIndx++
		}
	}

	return params, true
}

func (rd *routeDefault) Handlers(method string) []Handler {
	return rd.endpoints[method]
}

func (rd *routeDefault) URL() string {
	return rd.url
}

func (rd *routeDefault) Get(handlers ...Handler) IRoute {
	return rd.SetHandlers(http.MethodGet, handlers)
}
func (rd *routeDefault) Post(handlers ...Handler) IRoute {
	return rd.SetHandlers(http.MethodPost, handlers)
}
func (rd *routeDefault) Patch(handlers ...Handler) IRoute {
	return rd.SetHandlers(http.MethodPatch, handlers)
}
func (rd *routeDefault) Put(handlers ...Handler) IRoute {
	return rd.SetHandlers(http.MethodPut, handlers)
}
func (rd *routeDefault) Delete(handlers ...Handler) IRoute {
	return rd.SetHandlers(http.MethodDelete, handlers)
}
func (rd *routeDefault) Connect(handlers ...Handler) IRoute {
	return rd.SetHandlers(http.MethodConnect, handlers)
}
func (rd *routeDefault) Head(handlers ...Handler) IRoute {
	return rd.SetHandlers(http.MethodHead, handlers)
}
func (rd *routeDefault) Options(handlers ...Handler) IRoute {
	return rd.SetHandlers(http.MethodOptions, handlers)
}
func (rd *routeDefault) Trace(handlers ...Handler) IRoute {
	return rd.SetHandlers(http.MethodTrace, handlers)
}

func (rd *routeDefault) MultiMethod(methods []string, handlers ...Handler) IRoute {
	for _, method := range methods {
		rd.SetHandlers(method, handlers)
	}
	return rd
}

func skipToNextIndent(url string, urlLength int, currIndex int) int {
	for currIndex < urlLength && url[currIndex] != '/' {
		currIndex++
	}

	return currIndex
}

func extractToNextIndent(url string, urlLength int, currIndex int) string {
	start := currIndex
	end := skipToNextIndent(url, urlLength, currIndex)

	return url[start:end]
}
