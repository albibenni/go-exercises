# 03 - Live Code Review

## How to Structure Your Reasoning
1. State your review lens: correctness, reliability, readability, testability.
2. Identify issues by severity: bug > reliability > maintainability > style.
3. Propose concrete fixes with tradeoffs.
4. Mention tests that validate each fix.

## What Interviewers Look For
- Can you find real defects, not just style nits?
- Can you explain risk and impact?
- Can you prioritize changes under time pressure?
- Can you collaborate and reason aloud clearly?

## Review Checklist
- Input validation and edge cases
- Error propagation and context
- Retry/timeout/idempotency behavior
- Naming, function cohesion, boundary clarity
- Missing or weak tests

## Language to Use Live
- "I’d prioritize this because it can cause runtime failure/data corruption."
- "I’d keep this simple now; scale improvements can come later if needed."
- "I’d add a table-driven test for these edge cases."

## Red Flags in Code Review
- Vague comments without actionable fix
- Focusing on formatting before correctness
- No severity ordering
- Proposing refactors without test safety net
