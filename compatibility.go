package CloudPolicy

import (
	"encoding/json"
	"fmt"
)

type CommonValueBlock struct {
	value []string
}

var _ json.Marshaler = (*CommonValueBlock)(nil)
var _ json.Unmarshaler = (*CommonValueBlock)(nil)

func (t *CommonValueBlock) UnmarshalJSON(data []byte) error {
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

func (t *CommonValueBlock) MarshalJSON() ([]byte, error) {
	if len(t.value) == 1 {
		return json.Marshal(t.value[0])
	}
	return json.Marshal(t.value)
}

func (t *CommonValueBlock) SetString(s string) {
	t.value = []string{s}
}

func (t *CommonValueBlock) SetStrings(s ...string) {
	t.value = s
}

func (t *CommonValueBlock) Value() any {
	if len(t.value) == 1 {
		return t.value[0]
	}
	return t.value
}

type Principal struct {
	value map[PrincipalType]*CommonValueBlock
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

func (t *Principal) SetPrincipal(a PrincipalType, b *CommonValueBlock) {
	if t.value == nil {
		t.value = make(map[PrincipalType]*CommonValueBlock)
	}
	t.value = map[PrincipalType]*CommonValueBlock{a: b}
}

func (t *Principal) SetPrincipals(a map[PrincipalType]*CommonValueBlock) {
	t.value = a
}

func (t *Principal) Value() any {
	if t.all {
		return "*"
	}
	return t.value
}
