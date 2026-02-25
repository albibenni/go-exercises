# 02 - Concurrency

## Review Checklist
- Goroutine lifecycle ownership
- Channel direction and buffering choices
- `select` with timeout/cancellation
- `context.Context` propagation
- Mutex usage and lock scope
- Data races (`go test -race`)

## High-Yield Questions
- Channel vs mutex: when and why?
- How do goroutine leaks happen?
- How do you stop worker pools cleanly?
- What should be canceled with context?

## Practice Tasks
1. Build a worker pool with cancellation.
2. Add timeout handling to a channel receive path.
3. Introduce and then fix a race with shared state.

## Red Flags in Code Review
- Goroutines launched without shutdown path
- Blocking send/receive with no guarantee of counterpart
- Missing `ctx.Done()` handling
- Long critical sections under mutex
