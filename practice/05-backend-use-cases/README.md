# 05 - Backend Use Cases

## Review Checklist
- Validation at input boundaries
- Clear service/repository separation
- Transaction boundaries and rollback behavior
- Idempotency for retried requests
- Logging with useful context (not noise)

## Common Interview Use Cases
- Process order/payment with retries
- Ingest and deduplicate events
- Build simple REST endpoint with validation
- Batch process with partial failure handling

## Practice Tasks
1. Design a small service flow (handler -> service -> repo).
2. Add idempotency key handling to avoid duplicate processing.
3. Define failure scenarios and expected responses.

## Discussion Prompts
- What belongs in handler vs service vs repository?
- Where do you enforce validation and why?
- How do you make operations safe to retry?

## Red Flags in Code Review
- Business logic inside transport layer
- Missing timeout/retry strategy for external calls
- No handling for duplicate requests
