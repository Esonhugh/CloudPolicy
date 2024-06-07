## Cloud Policy

This library is the json base aws like policy document library.

Help you render the policy document in Go Sturct.

## Usage

```go
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
    var policy CloudPolicy.PolicyDocument
    err := json.Unmarshal([]byte(TestDocument_AliyunRestartEcsInstanceWithPrincipal), &policy)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(policy)
    fmt.Println(policy.Statement[0].Principal.Value())
    fmt.Println(policy.Statement[0].Condition)
    c := policy.Statement[0].Condition
    block := c["Bool"]["acs:MFAPresent"]
    fmt.Println(block.Value())
    fmt.Println(policy.Statement[0].Action)
    fmt.Println(policy.Statement[0].Effect)
    fmt.Println(policy.Statement[0].Resource)
    fmt.Println(policy.Version)
}
```