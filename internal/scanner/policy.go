package scanner

import "regexp"

const policyVersion = "0.1-rd.1"

type rule struct {
	id, severity, explanation, remediation string
	pattern                                *regexp.Regexp
}

var rules = []rule{
	{"secret.private-key", "critical", "Possible private-key header.", "Verify locally; if real, revoke or rotate and remove from release material.", regexp.MustCompile(`-----BEGIN (?:RSA |EC |OPENSSH |DSA |ENCRYPTED )?PRIVATE KEY-----`)},
	{"secret.github-token", "high", "Possible GitHub token.", "Verify locally; revoke or rotate any exposed credential.", regexp.MustCompile(`\b(?:gh[pousr]_[A-Za-z0-9]{36,255}|github_pat_[A-Za-z0-9_]{30,255})\b`)},
	{"secret.aws-access-id", "high", "Possible AWS access-key identifier (not proof of a usable credential).", "Verify ownership and exposure; rotate affected credentials if necessary.", regexp.MustCompile(`\b(?:AKIA|ASIA)[A-Z0-9]{16}\b`)},
	{"endpoint.internal", "medium", "Possible internal or local endpoint (preliminary).", "Confirm whether disclosure is intended and safe.", regexp.MustCompile(`(?i)https?://(?:localhost\b|127\.0\.0\.1\b|[a-z0-9.-]+\.(?:internal|local)\b)`)},
	{"ai.exposure-signal", "medium", "Possible AI configuration or instruction material (preliminary).", "Review for proprietary or sensitive information; this is not semantic analysis.", regexp.MustCompile(`(?i)\b(?:system_prompt|system_instruction|tool_schema|feature_flag)s?\b`)},
}
