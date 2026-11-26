package scanner

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func put(t *testing.T, root, name, value string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(root, name), []byte(value), 0600); err != nil {
		t.Fatal(err)
	}
}

func TestScanRedactsAndRequiresReview(t *testing.T) {
	d := t.TempDir()
	token := "gh" + "p_" + strings.Repeat("a", 36)
	put(t, d, "sample.txt", token)
	r, err := Scan(d)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Findings) != 1 || r.Findings[0].RuleID != "secret.github-token" || r.State != "VALIDATION_REQUIRED" || len(r.PendingChecks) < 4 {
		t.Fatalf("unexpected report: %+v", r)
	}
	b, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(b), token) {
		t.Fatal("secret material leaked")
	}
}

func TestManifestStableAndSensitiveToChange(t *testing.T) {
	d := t.TempDir()
	put(t, d, "a.txt", "first")
	a, err := Scan(d)
	if err != nil {
		t.Fatal(err)
	}
	b, err := Scan(d)
	if err != nil {
		t.Fatal(err)
	}
	if a.ScannedManifestSHA256 != b.ScannedManifestSHA256 || a.State != "VALIDATION_REQUIRED" {
		t.Fatal("unstable hash or approval implied")
	}
	put(t, d, "a.txt", "second")
	c, err := Scan(d)
	if err != nil {
		t.Fatal(err)
	}
	if a.ScannedManifestSHA256 == c.ScannedManifestSHA256 {
		t.Fatal("change undetected")
	}
}

func TestCoverageGaps(t *testing.T) {
	d := t.TempDir()
	put(t, d, "big.txt", strings.Repeat("a", int(maxFileBytes)+1))
	if err := os.Symlink("big.txt", filepath.Join(d, "link")); err != nil {
		t.Skip("symlinks unavailable")
	}
	r, err := Scan(d)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.CoverageGaps) != 3 || r.FilesScanned != 0 {
		t.Fatalf("unexpected coverage: %+v", r)
	}
	if _, err := Scan(filepath.Join(d, "link")); err == nil {
		t.Fatal("symlink root accepted")
	}
}

func TestPreliminaryRules(t *testing.T) {
	d := t.TempDir()
	put(t, d, "config.txt", "http://"+"localhost:8080\n"+"system_"+"prompt")
	r, err := Scan(d)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Findings) != 2 || r.Findings[0].Severity != "medium" || r.Findings[1].Severity != "medium" {
		t.Fatalf("unexpected findings: %+v", r.Findings)
	}
}
