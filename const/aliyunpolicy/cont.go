package aliyunpolicy

import cp "github.com/esonhugh/CloudPolicy"

const (
	PrincipalRAM       cp.PrincipalType = "RAM"
	PrincipalService   cp.PrincipalType = "Service"
	PrincipalFederated cp.PrincipalType = "Federated"
)

const (
	StringEquals              cp.ConditionOperation = "StringEquals"
	StringNotEquals           cp.ConditionOperation = "StringNotEquals"
	StringEqualsIgnoreCase    cp.ConditionOperation = "StringEqualsIgnoreCase"
	StringNotEqualsIgnoreCase cp.ConditionOperation = "StringNotEqualsIgnoreCase"
	StringLike                cp.ConditionOperation = "StringLike"
	StringNotLike             cp.ConditionOperation = "StringNotLike"

	NumericEquals            cp.ConditionOperation = "NumericEquals"
	NumericNotEquals         cp.ConditionOperation = "NumericNotEquals"
	NumericLessThan          cp.ConditionOperation = "NumericLessThan"
	NumericLessThanEquals    cp.ConditionOperation = "NumericLessThanEquals"
	NumericGreaterThan       cp.ConditionOperation = "NumericGreaterThan"
	NumericGreaterThanEquals cp.ConditionOperation = "NumericGreaterThanEquals"

	DateEquals            cp.ConditionOperation = "DateEquals"
	DateNotEquals         cp.ConditionOperation = "DateNotEquals"
	DateLessThan          cp.ConditionOperation = "DateLessThan"
	DateLessThanEquals    cp.ConditionOperation = "DateLessThanEquals"
	DateGreaterThan       cp.ConditionOperation = "DateGreaterThan"
	DateGreaterThanEquals cp.ConditionOperation = "DateGreaterThanEquals"

	Bool cp.ConditionOperation = "Bool"

	IPAddress    cp.ConditionOperation = "IpAddress"
	NotIPAddress cp.ConditionOperation = "NotIpAddress"
)

const (
	AcsCurrentTime     cp.ConditionKey = "acs:CurrentTime"
	AcsSecureTransport cp.ConditionKey = "acs:SecureTransport"
	AcsSourceIp        cp.ConditionKey = "acs:SourceIp"
	AcsMFAPresent      cp.ConditionKey = "acs:MFAPresent"
	AcsPrincipalArn    cp.ConditionKey = "acs:PrincipalArn"
	AcsPrincipalRDId   cp.ConditionKey = "acs:PrincipalRdId"
	AcsPrincipalRDPath cp.ConditionKey = "acs:PrincipalRdPath"
	AcsRequestTag      cp.ConditionKey = "acs:RequestTag"
	AcsResourceTag     cp.ConditionKey = "acs:ResourceTag"
)
