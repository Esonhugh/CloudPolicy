package CloudPolicy

import "encoding/json"

// PrincipalType is a string that represents the type of Principal. it should be like "AWS", "Federated", "Service", "CanonicalUser", "RAM", "qcs" .. or any other string.
type PrincipalType = string

const ()

// Principal is a map of PrincipalType and Value. or It will be "*"
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

// SetAllowAllPrincipal sets the Principal to "*"
func (t *Principal) SetAllowAllPrincipal() {
	t.all = true
}

// SetPrincipal sets the Principal to a single PrincipalType and Value
func (t *Principal) SetPrincipal(a PrincipalType, b *Value) {
	if t.value == nil {
		t.value = make(map[PrincipalType]*Value)
	}
	t.value = map[PrincipalType]*Value{a: b}
}

// SetPrincipals sets the Principal directly to a map of PrincipalType and Value
func (t *Principal) SetPrincipals(a map[PrincipalType]*Value) {
	t.value = a
}

// Value returns the value of the Principal object.
func (t *Principal) Value() any {
	if t.all {
		return "*"
	}
	return t.value
}
