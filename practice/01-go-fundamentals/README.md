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

Define:
Order struct (ID, Items, Status, Total)
OrderStore struct with map[string]Order
Processor interface with Process(order *Order) error
Implement:
AddOrder(order Order) and AddOrderPtr(order*Order) to show value vs pointer behavior.
GetOrInitOrders() []Order where the internal slice can be nil; handle safely.
A function that appends items to a passed slice and explain/observe append side effects (capacity reallocation vs shared backing array).
Use methods:
func (s OrderStore) Count() int (value receiver)
func (s *OrderStore) Save(o Order) (pointer receiver)
Then explain which methods are in the method set of OrderStore vs*OrderStore.
Interfaces:
Create type Logger interface { Log(msg string) }
Have ConsoleLogger implement it implicitly (no implements keyword).
Reliability:
In Process, use defer to log completion time.
Introduce a panic for corrupted input (nil item list), and recover in a top-level SafeProcess wrapper using defer + recover.
Errors:
Return wrapped errors with context, e.g.:
fmt.Errorf("validate order %s: %w", o.ID, err)
fmt.Errorf("save order %s: %w", o.ID, err)
