package rem

import (
	"regexp"
)

func NewVariableRoute(url string) *VariableRoute {
	route := VariableRoute{}
	route.initFromURL(cleanPath(url))
	return &route
}

// A route that can contain variables
type VariableRoute struct {
	BasicRoute
	// whether the route has any variables or not
	hasVariables    bool
}

func (vr *VariableRoute) initFromURL(url string) {
	vr.BasicRoute.url = url
	vr.BasicRoute.endpoints = map[string][]Handler{}

	hasMatched, err := regexp.Match(":\\w+", []byte(url))
	if err != nil {
		panic(err)
	}
	vr.hasVariables = hasMatched
}

func (vr *VariableRoute) skipToNextIndent(url *string, urlLength int, currIndex int) int {
	for currIndex < urlLength &&
		(*url)[currIndex] != '/' {
		currIndex++
	}

	return currIndex
}

func (vr *VariableRoute) extractToNextIndent(url *string, urlLength int, currIndex int) *string {
	out := ""

	for currIndex < urlLength &&
		(*url)[currIndex] != '/' {
		out += string((*url)[currIndex])
		currIndex++
	}

	return &out
}

func (vr *VariableRoute) Match(targetURL string) bool {
	if !vr.hasVariables {
		return vr.url == targetURL
	}

	originalURL := vr.BasicRoute.url
	oLen := len(originalURL)
	tLen := len(targetURL)

	oIndx := 0
	tIndx := 0
	for true {
		if originalURL[oIndx] == ':' {
			tIndx = vr.skipToNextIndent(&targetURL, tLen, tIndx)
			oIndx = vr.skipToNextIndent(&originalURL, oLen, oIndx)
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

func (vr *VariableRoute) extractURLParameters(targetURL string) *map[string]string {
	urlParameters := map[string]string{}

	originalURL := vr.BasicRoute.url
	oLen := len(originalURL)
	tLen := len(targetURL)

	oIndx := 0
	tIndx := 0
	for true {
		if originalURL[oIndx] == ':' {
			variable := vr.extractToNextIndent(&targetURL, tLen, tIndx)
			name     := vr.extractToNextIndent(&originalURL, oLen, oIndx + 1)
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