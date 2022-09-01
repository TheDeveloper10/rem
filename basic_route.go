package rem

import (
	"net/http"
	"regexp"
)

func NewBasicRoute(url string) *BasicRoute {
	route := BasicRoute{
		url: 	   cleanPath(url),
		endpoints: map[string][]Handler{},
	}
	route.initFromURL(url)

	return &route
}

type BasicRoute struct {
	url      string
	endpoints map[string][]Handler
	// whether the route has any variables or not
	hasVariables    bool
}

func (br *BasicRoute) initFromURL(url string) {
	br.url = cleanPath(url)
	br.endpoints = map[string][]Handler{}

	hasMatched, err := regexp.Match(":\\w+", []byte(url))
	if err != nil {
		panic(err)
	}
	br.hasVariables = hasMatched
}


func (br *BasicRoute) skipToNextIndent(url *string, urlLength int, currIndex int) int {
	for currIndex < urlLength &&
		(*url)[currIndex] != '/' {
		currIndex++
	}

	return currIndex
}

func (br *BasicRoute) extractToNextIndent(url *string, urlLength int, currIndex int) *string {
	out := ""

	for currIndex < urlLength &&
		(*url)[currIndex] != '/' {
		out += string((*url)[currIndex])
		currIndex++
	}

	return &out
}

// This method manages the incoming request
func (br *BasicRoute) handle(response IResponse, request IRequest) {
	handlers, ok := br.endpoints[request.GetMethod()]
	if !ok {
		response.Status(http.StatusMethodNotAllowed)
		return
	}

	for _, handler := range handlers {
		status := handler(response, request)
		if !status {
			return
		}
	}
}

func (br *BasicRoute) setHandlers(method string, handlers []Handler) IRoute {
	br.endpoints[method] = handlers
	return br
}

func (br *BasicRoute) Match(targetURL string) bool {
	if !br.hasVariables {
		return br.url == targetURL
	}

	originalURL := br.url
	oLen := len(originalURL)
	tLen := len(targetURL)

	oIndx := 0
	tIndx := 0
	for true {
		if originalURL[oIndx] == ':' {
			tIndx = br.skipToNextIndent(&targetURL, tLen, tIndx)
			oIndx = br.skipToNextIndent(&originalURL, oLen, oIndx)
		} else if originalURL[oIndx] != targetURL[tIndx] {
			return false
		}

		oIndx++
		tIndx++

		oOverTop := oIndx >= oLen
		tOverTop := tIndx >= tLen

		if oOverTop && tOverTop {
			return true
		} else if !oOverTop && !tOverTop {
			continue
		} else {
			return false
		}
	}

	return false
}

func (br *BasicRoute) extractURLParameters(targetURL string) *map[string]string {
	urlParameters := map[string]string{}

	originalURL := br.url
	oLen := len(originalURL)
	tLen := len(targetURL)

	oIndx := 0
	tIndx := 0
	for true {
		if originalURL[oIndx] == ':' {
			variable := br.extractToNextIndent(&targetURL, tLen, tIndx)
			name     := br.extractToNextIndent(&originalURL, oLen, oIndx + 1)
			tIndx = tIndx + len(*variable)
			oIndx = oIndx + len(*name) + 1
			urlParameters[*name] = *variable
		} else if originalURL[oIndx] != targetURL[tIndx] {
			return &urlParameters
		}

		oIndx++
		tIndx++

		oOverTop := oIndx >= oLen
		tOverTop := tIndx >= tLen

		if oOverTop && tOverTop {
			break
		} else if !oOverTop && !tOverTop {
			continue
		} else {
			break
		}
	}

	return &urlParameters
}

func (br *BasicRoute) GetRoute(handlers ...Handler)    IRoute { return br.setHandlers(http.MethodGet,    handlers) }
func (br *BasicRoute) PostRoute(handlers ...Handler)   IRoute { return br.setHandlers(http.MethodPost,   handlers) }
func (br *BasicRoute) PatchRoute(handlers ...Handler)  IRoute { return br.setHandlers(http.MethodPatch,  handlers) }
func (br *BasicRoute) PutRoute(handlers ...Handler)    IRoute { return br.setHandlers(http.MethodPut,    handlers) }
func (br *BasicRoute) DeleteRoute(handlers ...Handler) IRoute { return br.setHandlers(http.MethodDelete, handlers) }