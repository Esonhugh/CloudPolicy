package awspolicy

import (
	cp "github.com/esonhugh/CloudPolicy"
)

// Principal Types
const (
	AWS           cp.PrincipalType = "AWS"
	Federated     cp.PrincipalType = "Federated"
	Service       cp.PrincipalType = "Service"
	CanonicalUser cp.PrincipalType = "CanonicalUser"
)

// Condition Operations
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

	ArnEquals    cp.ConditionOperation = "ArnEquals"
	ArnLike      cp.ConditionOperation = "ArnLike"
	ArnNotEquals cp.ConditionOperation = "ArnNotEquals"
	ArnNotLike   cp.ConditionOperation = "ArnNotLike"

	IfExistsSuffix cp.ConditionOperation = "IfExists"

	Null cp.ConditionOperation = "Null"
)

// Global Condition Keys
const (
	// Principal
	AWSPrincipalArn              cp.ConditionKey = "aws:PrincipalArn"
	AWSPrincipalAccount          cp.ConditionKey = "aws:PrincipalAccount"
	AWSPrincipalOrgPath          cp.ConditionKey = "aws:PrincipalOrgPath"
	AWSPrincipalOrgID            cp.ConditionKey = "aws:PrincipalOrgID"
	AWSPrincipalTag              cp.ConditionKey = "aws:PrincipalTag"
	AWSPrincipalIsAWSService     cp.ConditionKey = "aws:PrincipalIsAWSService"
	AWSPrincipalServiceName      cp.ConditionKey = "aws:PrincipalServiceName"
	AWSPrincipalServiceNamesList cp.ConditionKey = "aws:PrincipalServiceNamesList"
	AWSPrincipalType             cp.ConditionKey = "PrincipalType"
	AWSUserID                    cp.ConditionKey = "aws:userid"
	AWSUsername                  cp.ConditionKey = "aws:username"

	// Role
	AWSFederatedProvider            cp.ConditionKey = "aws:FederatedProvider"
	AWSTokenIssueTime               cp.ConditionKey = "aws:TokenIssueTime"
	AWSMultiFactorAuthAge           cp.ConditionKey = "aws:MultiFactorAuthAge"
	AWSMultiFactorAuthPresent       cp.ConditionKey = "aws:MultiFactorAuthPresent"
	AWSEc2InstanceSourceVpc         cp.ConditionKey = "aws:Ec2InstanceSourceVpc"
	AWSEc2InstanceSourcePrivateIPv4 cp.ConditionKey = "aws:Ec2InstanceSourcePrivateIPv4"
	AWSSourceIdentity               cp.ConditionKey = "aws:SourceIdentity"
	Ec2RoleDelivery                 cp.ConditionKey = "ec2:RoleDelivery"
	Ec2SourceInstanceArn            cp.ConditionKey = "ec2:SourceInstanceArn"
	GlueRoleAssumedBy               cp.ConditionKey = "glue:RoleAssumedBy"
	GlueCredentialIssueService      cp.ConditionKey = "glue:CredentialIssueService"
	LambdaSourceFunctionArn         cp.ConditionKey = "lambda:SourceFunctionArn"
	SSMSourceInstanceArn            cp.ConditionKey = "ssm:SourceInstanceArn"
	IdentityStoreUserId             cp.ConditionKey = "identitystore:UserId"

	// Network
	AWSSourceIP    cp.ConditionKey = "aws:SourceIp"
	AWSSourceVPC   cp.ConditionKey = "aws:SourceVpc"
	AWSSourceVPCE  cp.ConditionKey = "aws:SourceVpce"
	AWSVPCSourceIP cp.ConditionKey = "aws:VpcSourceIp"

	// Properties of Resource
	AWSResourceAccount  cp.ConditionKey = "aws:ResourceAccount"
	AWSResourceOrgID    cp.ConditionKey = "aws:ResourceOrgID"
	AWSResourceOrgPaths cp.ConditionKey = "aws:ResourceOrgPaths"
	AWSResourceTag      cp.ConditionKey = "aws:ResourceTag"

	// Request
	AWSCalledVia       cp.ConditionKey = "aws:CalledVia"
	AWSCalledVidFirst  cp.ConditionKey = "aws:CalledViaFirst"
	AWSCalledViaLast   cp.ConditionKey = "aws:CalledViaLast"
	AWSViaAWSService   cp.ConditionKey = "aws:ViaAWSService"
	AWSCurrentTime     cp.ConditionKey = "aws:CurrentTime"
	AWSEpochTime       cp.ConditionKey = "aws:EpochTime"
	AWSReferer         cp.ConditionKey = "aws:Referer"
	AWSReqyestedRegion cp.ConditionKey = "aws:RequestedRegion"
	AWSRequestTag      cp.ConditionKey = "aws:RequestTag"
	AWSTagKeys         cp.ConditionKey = "aws:TagKeys"
	AWSSecureTransport cp.ConditionKey = "aws:SecureTransport"
	AWSSourceArn       cp.ConditionKey = "aws:SourceArn"
	AWSSourceAccount   cp.ConditionKey = "aws:SourceAccount"
	AWSSourceOrgID     cp.ConditionKey = "aws:SourceOrgID"
	AWSSourceOrgPaths  cp.ConditionKey = "aws:SourceOrgPaths"
	AWSUserAgent       cp.ConditionKey = "aws:UserAgent"
)
