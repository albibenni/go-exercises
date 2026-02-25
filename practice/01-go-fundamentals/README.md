# 01 - Go Fundamentals

## Review Checklist

- Value vs pointer semantics
- Slices and maps (`nil`, capacity, append side effects)
- Struct methods and method sets
- Interfaces and implicit implementation
- `defer`, `panic`, `recover`
- Error handling and wrapping (`%w`)

## High-Yield Questions

- When should a method have pointer receiver?
- Why can appending to a slice in a helper cause surprises?
- What is the zero value behavior for maps and slices?
- How do you preserve error context correctly?

## Practice Tasks

1. Implement a function that updates a struct both by value and pointer; explain the difference.
2. Write a helper that appends to a slice and verify when caller sees changes.
3. Refactor nested error handling using wrapped errors.

## Red Flags in Code Review

- Ignored returned errors
- Hidden mutation through shared slice backing array
- Overuse of `interface{}` / `any` without need
- Panic used for business logic
