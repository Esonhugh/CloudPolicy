package CloudPolicy

// Effect is the effect of a statement in a policy document. it should only be "Allow" or "Deny" or "allow" or "deny" things.
type Effect = string

const (
	EffectAllow  Effect = "Allow"
	EffectAllowl Effect = "allow"
	EffectDeny   Effect = "Deny"
	EffectDenyl  Effect = "deny"
)
