package CloudPolicy

type (
	// ConditionOperation is a string that represents the operation of the condition. it should be like "IpAddress", "Bool" .. or any other string.
	ConditionOperation = string

	// ConditionKey is a string that represents the key of the condition. it should be like "acs:MFAPresent", "acs:SourceIp" .. or any other string.
	ConditionKey = string
)

const (
	ConditionOperationIPAddress ConditionOperation = "IpAddress"
	ConditionOperationBool      ConditionOperation = "Bool"
)

// Condition is a map of ConditionOperation and ConditionValueList.
type Condition map[ConditionOperation]*ConditionValueList

// ConditionValueList is a map of ConditionKey and Value.
type ConditionValueList map[ConditionKey]*Value

// Add adds a new ConditionOperation and ConditionValueList to the Condition.
func (c *Condition) Add(key ConditionOperation, cl *ConditionValueList) *Condition {
	(*c)[key] = cl
	return c
}

// Del deletes a ConditionOperation from the Condition.
func (c *Condition) Del(key ConditionOperation) *Condition {
	delete(*c, key)
	return c
}

// NewSubCondition creates a new ConditionValueList. It's a helper function.
func NewSubCondition() *ConditionValueList {
	return &ConditionValueList{}
}

// Add adds a new ConditionKey and Value to the ConditionValueList.
func (c *ConditionValueList) Add(key ConditionKey, value *Value) *ConditionValueList {
	(*c)[key] = value
	return c
}

// Del deletes a ConditionKey from the ConditionValueList.
func (c *ConditionValueList) Del(key ConditionKey) *ConditionValueList {
	delete(*c, key)
	return c
}
