package CloudPolicy

import (
	"encoding/json"
	"fmt"
)

// Value is a string or an array of strings in order to compatible with AWS policy list and string type.
type Value struct {
	value []string
}

// NewValue creates a new Value object.
func NewValue(s ...string) *Value {
	return (&Value{}).Set(s...)
}

var _ json.Marshaler = (*Value)(nil)
var _ json.Unmarshaler = (*Value)(nil)

func (t *Value) UnmarshalJSON(data []byte) error {
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch v := raw.(type) {
	case string:
		t.value = []string{v}
	case []interface{}:
		for _, item := range v {
			str, ok := item.(string)
			if !ok {
				return fmt.Errorf("non-string element found in array")
			}
			t.value = append(t.value, str)
		}
	default:
		return fmt.Errorf("expected string or array of strings, got %T", v)
	}
	return nil
}

func (t *Value) MarshalJSON() ([]byte, error) {
	if len(t.value) == 1 {
		return json.Marshal(t.value[0])
	}
	return json.Marshal(t.value)
}

// Set sets the value of the Value object. if value has any element, it will be cleaned.
func (t *Value) Set(s ...string) *Value {
	t.value = s
	return t
}

// Value returns the value of the Value object.
// it will return the first element if the value has only one element.
// otherwise, it will return the whole array.
func (t *Value) Value() any {
	if len(t.value) == 1 {
		return t.value[0]
	}
	return t.value
}
