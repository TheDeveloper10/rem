package rem

import (
	"encoding/json"
	"io"
)

type JSONBodyProcessor struct{}

func (bp *JSONBodyProcessor) Serialize(body any) ([]byte, error) {
	return json.Marshal(body)
}

func (bp *JSONBodyProcessor) Parse(body io.Reader, out any) error {
	return json.NewDecoder(body).Decode(&out)
}
