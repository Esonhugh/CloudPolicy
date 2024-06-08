package CloudPolicy

// PolicyDocument is the top-level structure for a policy document.
type PolicyDocument struct {
	Version   string
	Statement []Statement
}

// Statement is a single statement in a policy document.
type Statement struct {
	Effect      Effect     `json:",omitempty"` // Allow or Deny
	Principal   *Principal `json:",omitempty"`
	NoPrincipal *Principal `json:",omitempty"`
	Action      *Value     `json:",omitempty"`
	NotAction   *Value     `json:",omitempty"`
	Resource    *Value     `json:",omitempty"`
	NotResource *Value     `json:",omitempty"`
	Condition   Condition  `json:",omitempty"`
	Sid         string     `json:",omitempty"`
}

// NewPolicy will create a new PolicyDocument object. This is a helper function.
func NewPolicy() *PolicyDocument {
	return &PolicyDocument{}
}
