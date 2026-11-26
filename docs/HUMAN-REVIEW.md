# Human review and Proof-of-Impact record (R&D template)

This template is intentionally blank. The deciding human, not software or an AI agent, writes the approval entry. Keep sensitive evidence in an access-controlled location, not a public repository. No approval transfers to a changed artifact, recipient, version, cost, or scope.

## Before work

- Human owner / authorized operator:
- Outcome, allowed actions and data, time and spending bounds:
- Acceptance criteria and review checkpoint:

## Evidence package (prepared for human review)

- Repository, base commit, proposed commit or exact diff SHA-256:
- Candidate build artifact SHA-256 (separate from scan manifest):
- Tool/policy versions, scan time, report SHA-256 and protected location:
- Findings, severity, evidence and disposition for each:
- Coverage gaps and unperformed mandatory checks:
- Test/build results and CI identifiers:
- Dependencies, license, version, provenance and rollback review:
- Remaining uncertainty, alternatives and consequences of delay:

## Decision gate — human-authored only

- Deciding human identity and authenticated approval-record location:
- Timestamp, exact scope and immutable candidate identifiers:
- Decision (approve / block / conditional), rationale and conditions:
- Current-session authorization reference where required by AltmanAI policy:

A conditional decision is not release permission. Missing evidence, unresolved conditions, changed identifiers, or an unavailable reviewer mean stop. This file is a template, not authentication, enforcement, or proof of approval. No software here verifies identity or executes a release.

## Closeout — human-authored acceptance

- Observed result / destination verification:
- Proof: evidence reviewed:
- Alignment: purpose and acceptance criteria:
- Integrity: limitations and unresolved items:
- Humanity: retained override and correction paths:
- Impact: observed measurable outcome:
- Human acceptance, correction, or unresolved effects:

Until acceptance is recorded, state remains VALIDATION_REQUIRED. Stop queued actions on override; inspect the destination before retrying any possibly completed action. Retain only authorized evidence for the human-defined period; permit export, correction and deletion subject to applicable obligations.
