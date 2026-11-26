package scanner

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
	"unicode/utf8"
)

const maxFileBytes int64 = 2 * 1024 * 1024
const maxTotalBytes int64 = 100 * 1024 * 1024
const maxEntries = 10000

// Scan examines a local, stable directory. It never publishes or approves.
// Do not use against a directory that untrusted processes can mutate concurrently.
func Scan(root string) (Report, error) {
	r := Report{
		SchemaVersion: "0.1", PolicyVersion: policyVersion,
		ScannedAtUTC: time.Now().UTC().Format(time.RFC3339Nano), State: "VALIDATION_REQUIRED",
		Findings: []Finding{}, CoverageGaps: []CoverageGap{},
		PendingChecks: []string{"dependency vulnerability review", "license/compliance review", "version and build provenance verification", "human review and impact validation"},
	}
	root = filepath.Clean(root)
	info, err := os.Lstat(root)
	if err != nil || !info.IsDir() {
		return r, fmt.Errorf("scan root must be a readable directory, not a symlink")
	}
	h := sha256.New()
	entries := 0
	gap := func(path, reason string) { r.CoverageGaps = append(r.CoverageGaps, CoverageGap{path, reason}) }
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, walkErr error) error {
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if rel == "." {
			return walkErr
		}
		entries++
		if entries > maxEntries {
			gap(rel, "entry budget exhausted; remaining entries not scanned")
			return fs.SkipAll
		}
		if !utf8.ValidString(rel) {
			gap("[non-UTF-8 path]", "path cannot be represented losslessly in report")
			if d != nil && d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}
		if walkErr != nil {
			gap(rel, "cannot access entry")
			return nil
		}
		if d.Name() == ".git" {
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if !d.Type().IsRegular() {
			gap(rel, "non-regular file not scanned")
			return nil
		}
		if r.BytesScanned >= maxTotalBytes {
			gap(rel, "byte budget exhausted; remaining entries not scanned")
			return fs.SkipAll
		}
		info, err := d.Info()
		if err != nil {
			gap(rel, "cannot stat entry")
			return nil
		}
		if info.Size() > maxFileBytes || info.Size() > maxTotalBytes-r.BytesScanned {
			gap(rel, "file exceeds remaining scan budget")
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			gap(rel, "cannot open entry")
			return nil
		}
		limit := min(maxFileBytes, maxTotalBytes-r.BytesScanned)
		data, readErr := io.ReadAll(io.LimitReader(f, limit+1))
		closeErr := f.Close()
		if readErr != nil || closeErr != nil || int64(len(data)) > limit {
			gap(rel, "read failed or budget exceeded")
			return nil
		}
		sum := sha256.Sum256(data)
		// JSON-encoded [path, hash] records followed by LF avoid ambiguous concatenation.
		entry, err := json.Marshal([2]string{rel, hex.EncodeToString(sum[:])})
		if err != nil {
			return err
		}
		_, _ = h.Write(entry)
		_, _ = h.Write([]byte{'\n'})
		r.FilesScanned++
		r.BytesScanned += int64(len(data))
		for _, rule := range rules {
			if rule.pattern.Match(data) {
				r.Findings = append(r.Findings, Finding{rule.id, rule.severity, rel, rule.explanation, rule.remediation})
			}
		}
		return nil
	})
	if err != nil {
		return r, fmt.Errorf("directory traversal failed")
	}
	r.ScannedManifestSHA256 = hex.EncodeToString(h.Sum(nil))
	if r.FilesScanned == 0 {
		gap(".", "no regular files scanned")
	}
	return r, nil
}
