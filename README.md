# ReleaseShield by AltmanAI

**Status: early development / v0.1 R&D candidate — not a production release gate.**

ReleaseShield is a human-first release-integrity project for inspecting candidate software, surfacing risk and collecting evidence for accountable human decisions. AI may assist analysis; it does not authorize releases.

Humanity leads. Intelligence follows.

## Workflow

**Scan → Flag → Human Review → Prove.** Human authorization defines scope before scanning. The human reviews findings and evidence before any commitment or release, and validates the actual outcome at closeout. Changed scope or candidate identifiers require fresh review. See [the human review record](docs/HUMAN-REVIEW.md).

## Development capability and v0.1 target

| Area | This R&D candidate |
| --- | --- |
| Secret/token exposure | Limited pattern rules for private-key headers, GitHub tokens and AWS access-key identifiers; not exhaustive |
| Dependency vulnerabilities | Required before release; not implemented |
| License/compliance | Required before release; not implemented |
| Internal endpoints | Preliminary pattern detection |
| AI exposure signals | Preliminary pattern detection; no semantic analysis |
| Evidence | JSON with policy/schema version, time, severity, relative path, remediation, coverage gaps and scanned-content manifest hash |
| Human authorization | Separate human-authored review template; no identity verification or automated release enforcement |
| CI | Proposed formatting, vet, race tests and build checks; not deployment |

## Local development

Requires Go 1.23 or newer to build this candidate. Use a currently supported and patched Go toolchain for development and release; the module minimum is not a security recommendation.

```sh
go vet ./...
go test -race ./...
go build -o /tmp/releaseshield ./cmd/releaseshield
/tmp/releaseshield scan ./candidate > /tmp/releaseshield-report.json
```

Run on a stable, authorized local directory; keep the report outside the scanned directory. The CLI performs no network calls, uses only the Go standard library and does not publish or deploy. Exit status `3` means a report was produced and human validation is required (including when no rules matched). Status `1` indicates a scan/output error; `2` indicates invalid usage. No status grants release permission.

## Boundaries and evidence

The scanner excludes `.git` metadata, does not follow symlinks and reports skipped non-regular, inaccessible or oversized files, non-UTF-8 paths and exhausted budgets. Limits: 2 MiB per file, 100 MiB of successfully scanned bytes and 10,000 directory entries. It scans regular file bytes, not archives' extracted contents or Git history. Use a stable snapshot that untrusted processes cannot mutate; concurrent path replacement is not defended against. Findings omit matched content; relative paths may themselves be sensitive. Do not publish raw reports without review.

The manifest SHA-256 covers JSON-encoded relative path and file-content SHA-256 pairs, each followed by LF, in traversal order. It binds only successfully scanned file contents and paths; it is not a build-artifact hash, signed attestation, full filesystem snapshot or tamper-proof ledger. Zero findings does not mean safe. False positives and false negatives are expected. Dependency, license, version and provenance checks remain explicitly pending.

## R&D milestones

1. Review this baseline and validate local/CI tests on supported toolchains.
2. Integrate vetted dependency and license analysis with versioned policies and explicit error/coverage handling; test on labeled synthetic fixtures.
3. Bind reviewed candidate commits, build artifacts and report hashes to authenticated human decisions; validate pause, scope changes and rollback.
4. Pilot on authorized repositories; measure seeded detection coverage, false positives, review time and audit completeness. Require independent human acceptance before any release claim.

## Security and governance

Report potential vulnerabilities privately to security@altmanai.tech; do not include live secrets in public issues. General inquiries: support@altmanai.tech.

Related public references: [Masterplan](https://github.com/altmanAI/altmanai-masterplan), [P.A.I.H.I. framework](https://github.com/altmanAI/paihi-framework), [Human-First AI](https://github.com/altmanAI/human-first-ai). These are contextual references, not claims of implemented integrations.

Maintained by AltmanAI, Altman Family Group, LLC. Existing license: [Apache-2.0](LICENSE).
