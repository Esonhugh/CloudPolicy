package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/esonhugh/CloudPolicy"
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
	var policy *CloudPolicy.PolicyDocument
	err := json.Unmarshal([]byte(TestDocument_AliyunRestartEcsInstanceWithPrincipal), &policy)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(policy)
	fmt.Println(policy.Statement[0].Principal.Value())
	fmt.Println(policy.Statement[0].Condition)
	(*policy.Statement[0].Condition["Bool"])["acs:MFAPresent"].SetString("false")
	a := CloudPolicy.NewConditionValueList()
	a.Add("acs:SourceIp", CloudPolicy.NewCommonValueBlock().SetString("127.0.0.1"))
	policy.Statement[0].Condition[CloudPolicy.ConditionOperationIPAddress] = a
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

	newP := CloudPolicy.PolicyDocument{
		Version: "1",
		Statement: []CloudPolicy.Statement{
			{
				Effect:   "Allow",
				Action:   CloudPolicy.NewCommonValueBlock().SetString("ecs:RebootInstance"),
				Resource: CloudPolicy.NewCommonValueBlock().SetString("*"),
				Condition: CloudPolicy.Condition{
					CloudPolicy.ConditionOperationBool: CloudPolicy.NewConditionValueList().
						Add("acs:MFAPresent", CloudPolicy.NewCommonValueBlock().SetString("true")).
						Add("acs:SourceIp", CloudPolicy.NewCommonValueBlock().SetString("false")),
				},
			},
		},
	}
	fmt.Println(newP)
}
