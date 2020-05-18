package json

import "encoding/json"

func Jsonify(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
