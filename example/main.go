package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	cp "github.com/esonhugh/CloudPolicy"
	"github.com/esonhugh/CloudPolicy/const/awspolicy"
)

// Example of decoding a policy from a JSON string and changing its value.
func DecodeFromJsonExample() {
	fmt.Println("DecodeFromJsonExample")
	const TestDocument_AliyunRestartEcsInstanceWithPrincipal = `
{
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
	var policy *cp.PolicyDocument
	err := json.Unmarshal([]byte(TestDocument_AliyunRestartEcsInstanceWithPrincipal), &policy)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(policy)
	fmt.Println(policy.Statement[0].Principal.Value())
	fmt.Println(policy.Statement[0].Condition)
	fmt.Println(policy.Statement[0].Action)
	fmt.Println(policy.Statement[0].Effect)
	fmt.Println(policy.Statement[0].Resource)
	fmt.Println(policy.Version)

	// Change the value of the condition
	(*policy.Statement[0].Condition["Bool"])["acs:MFAPresent"].Set("false")
	// Add a new kind of condition in policy
	policy.Statement[0].Condition.Add("DATA", &cp.ConditionValueList{
		"tagName": cp.NewValue("helper"),
	})
	// Delete the condition
	delete(policy.Statement[0].Condition, "DATA")

	// Add a new condition 2 (another flavor)
	a := cp.NewSubCondition()
	a.Add("acs:SourceIp", cp.NewValue("127.0.0.1"))
	policy.Statement[0].Condition[cp.ConditionOperationIPAddress] = a

	// Print the policy after changed
	fmt.Println("After change")
	res, e := json.Marshal(policy)
	if e != nil {
		fmt.Println(e)
		os.Exit(-1)
	}
	fmt.Println(string(res))
	res, e = json.Marshal(policy.Statement[0].Condition)
	if e != nil {
		fmt.Println(e)
		os.Exit(-1)
	}
	fmt.Println(string(res))
}

// Example of dynamic construction of a new policy
func ConstructNewPolicyExample() {
	newPolicy := cp.PolicyDocument{
		Version: "1",
		Statement: []cp.Statement{
			// First statement create via directly setting the values
			{
				Effect:   cp.EffectAllow,
				Action:   cp.NewValue().Set("ecs:RebootInstance"),
				Resource: cp.NewValue("*"),
				Condition: cp.Condition{
					"Bool": &cp.ConditionValueList{
						awspolicy.AWSRequestTag: cp.NewValue("true"),
						awspolicy.AWSSourceIP:   cp.NewValue("127.0.0.1"),
					},
				},
			},
			// Second statement create via setting the values one by one with chained helper function
			{
				Effect: cp.EffectAllow,
				Action: cp.NewValue(
					"oss:GetObject",
					"oss:ListObjects",
				),
				Resource: cp.NewValue("arn::uuid:bucket/helper"),
				Condition: *cp.NewCondition().
					Add("ops", cp.NewSubCondition().
						Add("KeyOfAAAA", cp.NewValue("127.0.0.1")).
						Add("KeyOfBBBB", cp.NewValue("false"))).
					Add("ops2", cp.NewSubCondition().
						Add("KeyOfCCCC", cp.NewValue("true")).
						Add("KeyOfDDDD", cp.NewValue("false"))),
			},
		},
	}
	fmt.Println(newPolicy)
	res, e := json.Marshal(newPolicy)
	if e != nil {
		fmt.Println(e)
		os.Exit(-2)
	}
	fmt.Println(string(res))
}

func main() {
	ConstructNewPolicyExample()
	DecodeFromJsonExample()
}
