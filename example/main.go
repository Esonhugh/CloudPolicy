package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	p "github.com/esonhugh/CloudPolicy"
)

func main() {
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
	var policy *p.PolicyDocument
	err := json.Unmarshal([]byte(TestDocument_AliyunRestartEcsInstanceWithPrincipal), &policy)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(policy)
	fmt.Println(policy.Statement[0].Principal.Value())
	fmt.Println(policy.Statement[0].Condition)
	(*policy.Statement[0].Condition["Bool"])["acs:MFAPresent"].Set("false")
	policy.Statement[0].Condition.Add("DATA", &p.ConditionValueList{
		"tagName": p.NewValue("helper"),
	})
	delete(policy.Statement[0].Condition, "DATA")

	a := p.NewSubCondition()
	a.Add("acs:SourceIp", p.NewValue("127.0.0.1"))
	policy.Statement[0].Condition[p.ConditionOperationIPAddress] = a

	fmt.Println(policy.Statement[0].Action)
	fmt.Println(policy.Statement[0].Effect)
	fmt.Println(policy.Statement[0].Resource)
	fmt.Println(policy.Version)

	fmt.Println("After change")
	res, e := json.Marshal(policy)
	if e != nil {
		fmt.Println(e)
		os.Exit(-1)
	}
	fmt.Println(string(res))
	fmt.Println(policy.Statement[0].Principal.Value())

	newP := p.PolicyDocument{
		Version: "1",
		Statement: []p.Statement{
			{
				Effect:   "Allow",
				Action:   p.NewValue().Set("ecs:RebootInstance"),
				Resource: p.NewValue().Set("*"),
				Condition: p.Condition{
					"Bool": &p.ConditionValueList{
						"acs:MFAPresent": p.NewValue("true"),
						"acs:SourceIp":   p.NewValue("false"),
					},
				},
			},
			{
				Effect: p.EffectAllow,
				Action: p.NewValue().Set([]string{
					"oss:GetObject",
					"oss:ListObjects",
				}),
				Resource: p.NewValue().Set("arn::uuid:bucket/helper"),
				Condition: p.Condition{
					"ops": p.NewSubCondition().
						Add("aaa", p.NewValue("falsex")),
					"ops2": p.NewSubCondition().
						Add("SourceIPs", p.NewValue("127.0.0.1")),
					"ops3": p.NewSubCondition().Add("12", p.NewValue("122", "2333")),
				},
			},
		},
	}
	fmt.Println(newP)
	res, e = json.Marshal(newP)
	if e != nil {
		fmt.Println(e)
		os.Exit(-2)
	}
	fmt.Println(string(res))
}
