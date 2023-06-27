package tjson

import "encoding/json"

type Json interface {
	any
}

type Jsonable interface {
	ToJson() string
}

func ToJson[T Json](v T) string {
	j, _ := json.Marshal(v)
	return string(j)
}

func ToJsonPretty[T Json](v T) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

func FromJson[T Json](s string) T {
	var p T
	json.Unmarshal([]byte(s), &p)
	return p
}
