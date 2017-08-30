package util

import (
	"bytes"
	"encoding/json"
)

func Json(obj interface{}) string {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(obj)

	return buf.String()
}
