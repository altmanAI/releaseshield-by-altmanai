# ReleaseShield by AltmanAI

ReleaseShield by AltmanAI is a human-first release integrity system that tracks builds, approvals, hashes, changelogs, and Proof-of-Impact records. It gives teams clearer accountability and verifiable trust from first commit to public release. AI surfaces risk and context. Humans make the final decision. Proof is generated at every step.

## What is ReleaseShield?

ReleaseShield helps teams ship software with confidence by scanning for risk before release, requiring human review of flagged issues, and generating structured proof reports. It is not another scanner — it is an orchestration + policy + proof layer that keeps the human accountable.

**Core principle**: Humanity leads. Intelligence follows.

## v0.1 Scope (Honest)

ReleaseShield v0.1 focuses on five scan categories:

- **Secrets / token exposure** (mandatory)
- **Vulnerable dependencies** (mandatory)
- **License / compliance issues** (mandatory)
- **Internal URL / endpoint exposure** (preliminary — pattern-based)
- **AI-specific exposure signals** (preliminary — prompts, system instructions, tool schemas, feature flags — pattern-based only)

**Important**: Categories marked “preliminary” use lightweight detection and explicitly require human verification. v0.1 does not claim deep semantic or model-level analysis.

## How It Works

1. **Scan** — Run the CLI against a folder or repository
2. **Flag** — Findings are categorized with severity and context
3. **Human Review** — Required review step with decision (approve / block / conditional) and notes
4. **Proof Report** — Structured JSON + human-readable Markdown containing severity, file path, explanation, remediation, timestamp, artifact hash, and reviewer decision

## Key Features (v0.1)

- Local CLI that runs against any folder or repo
- JSON and Markdown report output
- Pass/fail status with clear severity levels
- Mandatory human review workflow
- GitHub Action integration (stub ready)
- Proof-grade reports with audit trail

## CLI Usage (Planned)

```bash
releaseshield scan ./my-release \
  --format json,md \
  --output ./reports/ \
  --severity medium,high,critical
