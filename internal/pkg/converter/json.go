package converter

import (
	"encoding/json"
)

// ToJSONString ...
func ToJSONString(v interface{}) (string, error) {
	bytesData, err := json.Marshal(v)

	return string(bytesData), err
}
