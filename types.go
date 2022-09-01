package rem

type KeyValues map[string][]string

func (kv KeyValues) Get(key string) (string, bool) {
	vals, ok := kv[key]
	if ok {
		if len(vals) > 0 {
			return vals[0], true
		} else {
			return "", true
		}
	} else {
		return "", false
	}
}



type KeyValue map[string]string

func (kv KeyValue) Get(key string) (string, bool) {
	val, ok := kv[key]
	if ok {
		return val, true
	} else {
		return "", false
	}
}