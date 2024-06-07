package CloudPolicy

import (
	"encoding/json"
	"fmt"
)

type Value struct {
	value []string
}

func NewValue(s ...string) *Value {
	return (&Value{}).Set(s)
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

func (t *Value) Set(s any) *Value {
	if s == nil {
		return t
	}

	switch v := s.(type) {
	case string:
		t.value = []string{v}
	case []string:
		t.value = v
	default:
		panic(fmt.Errorf("expected string or array of strings, got %T", v))
	}
	return t
}

func (t *Value) Value() any {
	if len(t.value) == 1 {
		return t.value[0]
	}
	return t.value
}

type Principal struct {
	value map[PrincipalType]*Value
	all   bool
}

var _ json.Marshaler = (*Principal)(nil)
var _ json.Unmarshaler = (*Principal)(nil)

func (t *Principal) UnmarshalJSON(data []byte) error {
	if string(data) == "\"*\"" {
		t.all = true
		return nil
	}
	t.all = false
	return json.Unmarshal(data, &t.value)
}

func (t *Principal) MarshalJSON() ([]byte, error) {
	if t.all {
		return json.Marshal("*")
	}
	return json.Marshal(t.value)
}

func (t *Principal) SetAllowAllPrincipal() {
	t.all = true
}

func (t *Principal) SetPrincipal(a PrincipalType, b *Value) {
	if t.value == nil {
		t.value = make(map[PrincipalType]*Value)
	}
	t.value = map[PrincipalType]*Value{a: b}
}

func (t *Principal) SetPrincipals(a map[PrincipalType]*Value) {
	t.value = a
}

func (t *Principal) Value() any {
	if t.all {
		return "*"
	}
	return t.value
}
