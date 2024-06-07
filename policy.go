package CloudPolicy

type PolicyDocument struct {
	Version   string
	Statement []Statement
}

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

type Effect = string

const (
	EffectAllow  Effect = "Allow"
	EffectAllowl Effect = "allow"
	EffectDeny   Effect = "Deny"
	EffectDenyl  Effect = "deny"
)

type PrincipalType = string

const (
	PrincipalTypeAWS           PrincipalType = "AWS"
	PrincipalTypeFederated     PrincipalType = "Federated"
	PrincipalTypeService       PrincipalType = "Service"
	PrincipalTypeCanonicalUser PrincipalType = "CanonicalUser"
	PrincipalTypeRAM           PrincipalType = "RAM"
	PrincipalTypeQCS           PrincipalType = "qcs"
)

type (
	ConditionOperation = string
	ConditionKey       = string
)

const (
	ConditionOperationIPAddress ConditionOperation = "IpAddress"
	ConditionOperationBool      ConditionOperation = "Bool"
)

type Condition map[ConditionOperation]*ConditionValueList

type ConditionValueList map[ConditionKey]*Value

func (c *Condition) Add(key ConditionOperation, cl *ConditionValueList) *Condition {
	(*c)[key] = cl
	return c
}

func NewSubCondition() *ConditionValueList {
	return &ConditionValueList{}
}

func (c *ConditionValueList) Add(key ConditionKey, value *Value) *ConditionValueList {
	(*c)[key] = value
	return c
}
