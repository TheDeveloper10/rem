package rem

type MapStringsReadOnly map[string][]string
type MapStringReadOnly map[string]string

func (mssro *MapStringsReadOnly) Get(key string) string {
	// like url.Values Get

	v, ok := (*mssro)[key]
	if !ok {
		return ""
	}

	return v[0]
}

func (msro *MapStringReadOnly) Get(key string) string {
	// like url.Values Get

	v, ok := (*msro)[key]
	if !ok {
		return ""
	}

	return v
}
