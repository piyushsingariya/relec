package relec

import (
	"encoding/json"
	"fmt"
)

func IsJSON[T ~string | []byte](val T) bool {
	var js json.RawMessage

	var anyVal any = val
	switch anyVal := anyVal.(type) {
	case []byte:
		return json.Unmarshal(anyVal, &js) == nil
	case string:
		return json.Unmarshal([]byte(anyVal), &js) == nil
	}

	return false
}

type JSONString struct {
	content string
}

func NewJSON(content string) JSONString {
	return JSONString{
		content: content,
	}
}

func (j *JSONString) UnmarshalJSON(data []byte) error {
	j.content = string(data)
	return nil
}

func (j *JSONString) MarshalJSON() ([]byte, error) {
	return []byte(j.content), nil
}

func (j *JSONString) Unmarshal(to any) error {
	if !IsPointer(to) {
		return fmt.Errorf("pass pointer to a variable: passed %T", to)
	}

	return json.Unmarshal([]byte(j.content), to)
}

func (j *JSONString) Any() (map[string]any, error) {
	to := make(map[string]any)

	return to, json.Unmarshal([]byte(j.content), &to)
}

func (j *JSONString) String() string {
	return j.content
}
