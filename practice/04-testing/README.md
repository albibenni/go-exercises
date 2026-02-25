# 04 - Testing

## Review Checklist
- Table-driven tests for behavior matrix
- Subtests with `t.Run`
- Deterministic tests (avoid flaky timing)
- Clear test names and arrange/act/assert structure
- Race detector and coverage awareness

## High-Yield Questions
- Why table-driven tests in Go?
- When should you use integration vs unit tests?
- How do you test timeout/cancellation paths?

## Practice Tasks
1. Convert one imperative test into table-driven style.
2. Add tests for edge cases: empty input, nil, duplicate, timeout.
3. Run race detector and fix one issue.

## Commands
- `go test ./...`
- `go test -race ./...`
- `go test -cover ./...`

## Red Flags in Code Review
- Tests asserting implementation details, not behavior
- Missing negative-path tests
- Sleep-based timing tests without bounds
