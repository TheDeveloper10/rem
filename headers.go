package rem

type Headers map[string][]string

func (h Headers) Get(key string) (*string, bool) {
	header, ok := h[key]
	if ok {
		res := ""
		if len(header) > 0 {
			res = header[0]
		}

		return &res, true
	} else {
		return nil, false
	}
}