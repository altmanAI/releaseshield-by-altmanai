package scanner

// Finding contains a rule reference and location, never matched secret material.
type Finding struct {
	RuleID      string `json:"rule_id"`
	Severity    string `json:"severity"`
	Path        string `json:"path"`
	Explanation string `json:"explanation"`
	Remediation string `json:"remediation"`
}

type CoverageGap struct {
	Path   string `json:"path"`
	Reason string `json:"reason"`
}

type Report struct {
	SchemaVersion         string        `json:"schema_version"`
	PolicyVersion         string        `json:"policy_version"`
	ScannedAtUTC          string        `json:"scanned_at_utc"`
	State                 string        `json:"state"`
	FilesScanned          int           `json:"files_scanned"`
	BytesScanned          int64         `json:"bytes_scanned"`
	ScannedManifestSHA256 string        `json:"scanned_manifest_sha256"`
	Findings              []Finding     `json:"findings"`
	CoverageGaps          []CoverageGap `json:"coverage_gaps"`
	PendingChecks         []string      `json:"pending_checks"`
}
