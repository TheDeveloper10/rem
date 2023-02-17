package rem

type KeyValues map[string][]string

func (kv KeyValues) Get(key string) string {
	vals, ok := kv[key]
	if ok {
		if len(vals) > 0 {
			return vals[0]
		}
	}
	return ""
}

func (kv KeyValues) set(key string, value []string) bool {
	kv[key] = value
	return true
}



type KeyValue map[string]string

func (kv KeyValue) Get(key string) string {
	val, ok := kv[key]
	if ok {
		return val
	}
	return ""
}
