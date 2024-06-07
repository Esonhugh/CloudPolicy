package CloudPolicy

import (
	"encoding/json"
	"testing"
)

func TestPolicies(t *testing.T) {
	t.Run("AliyunRestartEcsInstance", Test_AliyunRestartEcsInstance)
	t.Run("AliyunRestartEcsInstanceWithPrincipal", Test_AliyunRestartEcsInstanceWithPrincipal)
	t.Run("AWS_Principal", Test_AWS_Principal)
	t.Run("AliyunCreateSnapShot", Test_AliyunCreateSnapShot)
}

func CreatePolicyDocument(t *testing.T, policyString string) (res *PolicyDocument) {
	e := json.Unmarshal([]byte(policyString), &res)
	if e != nil {
		t.Error(e)
		t.Fail()
	}
	return res
}

func CreateJsonPolicy(t *testing.T, policy *PolicyDocument) (res []byte) {
	res, e := json.Marshal(policy)
	if e != nil {
		t.Error(e)
		t.Fail()
	}
	return res
}

const TestDocument_AliyunRestartEcsInstance = `{
  "Statement": [
    {
      "Action": "ecs:RebootInstance",
      "Effect": "Allow",
      "Resource": "*",
      "Condition": {
        "Bool": {
          "acs:MFAPresent": "true"
        }
      }
    }
  ],
  "Version": "1"
}`

func Test_AliyunRestartEcsInstance(t *testing.T) {
	policy := CreatePolicyDocument(t, TestDocument_AliyunRestartEcsInstance)
	t.Log(policy)
}

const TestDocument_AliyunRestartEcsInstanceWithPrincipal = `{
  "Statement": [
    {
      "Action": "ecs:RebootInstance",
      "Effect": "Allow",
      "Resource": "*",
      "Principal": { "AWS": "arn:aws:iam::123456789012:root" },
      "Condition": {
        "Bool": {
          "acs:MFAPresent": "true"
        }
      }
    }
  ],
  "Version": "1"
}`

func Test_AliyunRestartEcsInstanceWithPrincipal(t *testing.T) {
	policy := CreatePolicyDocument(t, TestDocument_AliyunRestartEcsInstanceWithPrincipal)
	t.Log(policy)
	str := CreateJsonPolicy(t, policy)
	t.Log(string(str))
}

const TestDocument_AWS_Principal = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "UsePrincipalArnInsteadOfNotPrincipalWithDeny",
      "Effect": "Deny",
      "Action": "s3:*",
      "Principal": "*",
      "Resource": [
        "arn:aws:s3:::BUCKETNAME/*",
        "arn:aws:s3:::BUCKETNAME"
      ],
      "Condition": {
        "ArnNotEquals": {
          "aws:PrincipalArn": "arn:aws:iam::444455556666:user/user-name"
        }
      }
    }
  ]
}`

func Test_AWS_Principal(t *testing.T) {
	policy := CreatePolicyDocument(t, TestDocument_AWS_Principal)
	t.Log(policy)
}

const TestDocument_AliyunCreateSnapShot = `{
  "Statement": [
    {
      "Action": "ecs:*",
      "Effect": "Allow",
      "Resource": [
        "acs:ecs:*:*:instance/i-bp14l21sgragk0sp****"
      ]
    },
    {
      "Action": "ecs:CreateSnapshot",
      "Effect": "Allow",
      "Resource": [
        "acs:ecs:*:*:disk/d-bp1hek8q2l5tqzde****",
        "acs:ecs:*:*:snapshot/*"
      ]
    },
    {
      "Action": [
        "ecs:Describe*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ],
  "Version": "1"
}`

func Test_AliyunCreateSnapShot(t *testing.T) {
	policy := CreatePolicyDocument(t, TestDocument_AliyunCreateSnapShot)
	t.Log(policy)
	str := CreateJsonPolicy(t, policy)
	t.Log(string(str))
}
