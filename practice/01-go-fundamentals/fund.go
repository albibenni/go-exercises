package main

import (
	"fmt"
	"time"
)

type Item struct {
	Name string
}

type Order struct {
	ID    string
	Total int
	Items []Item
}

type OrderStore struct {
	Orders map[string]Order
}

type ConsoleLogger struct{}

type Processor interface {
	Process(order *Order) error
}

type Logger interface {
	Log(msg string)
}

type ProcessorImpl struct {
	logger Logger
}

func (s *OrderStore) AddOrder(order Order) {
	if s.Orders == nil {
		s.Orders = map[string]Order{}
	}
	s.Orders[order.ID] = order
}
func (s *OrderStore) AddOrderPtr(order *Order) {
	if s.Orders == nil {
		s.Orders = map[string]Order{}
	}
	order.Total += 10
	s.Orders[order.ID] = *order
}

func (s *OrderStore) GetOrInitOrders() []Order {
	if s == nil {
		return nil
	}
	if s.Orders == nil {
		s.Orders = map[string]Order{}
	}
	orders := make([]Order, 0, len(s.Orders))
	for _, order := range s.Orders {
		orders = append(orders, order)
	}
	return orders
}

func (c ConsoleLogger) Log(msg string) {
	fmt.Println("[LOG]", msg)
}

func UseLogger(l Logger) {
	l.Log("order processed")
}

func AppendItem(items []Item, it Item) []Item {
	return append(items, it)
}

func (s OrderStore) Count() int {
	return len(s.Orders)
}
func (s *OrderStore) Save(o Order) {
	s.AddOrderPtr(&o)
}

func (p ProcessorImpl) Process(order *Order) error {
	start := time.Now()
	defer func() {
		p.logger.Log(fmt.Sprintf("Process took %s", time.Since(start)))
	}()

	if order == nil || order.Items == nil {
		panic("corrupted order: nil order or nil items")
	}

	return nil
}

func SafeProcess(p Processor, order *Order) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("safe process recovered panic: %v", r)
		}
	}()
	return p.Process(order)
}

func main() {
	store := &OrderStore{Orders: map[string]Order{}}
	o1 := Order{ID: "A", Total: 100}
	store.AddOrder(o1)
	fmt.Println(o1.Total) // 100 (unchanged)

	o2 := Order{ID: "B", Total: 100}
	store.AddOrderPtr(&o2)
	fmt.Println(o2.Total) // 110 (changed)

	var nilOrders []Order
	fmt.Printf("nilOrders: len=%d cap=%d isNil=%v\n", len(nilOrders), cap(nilOrders), nilOrders == nil)
	nilOrders = append(nilOrders, Order{ID: "C", Total: 90})
	fmt.Printf("after append nilOrders: len=%d cap=%d isNil=%v\n", len(nilOrders), cap(nilOrders), nilOrders == nil)

	base := make([]Item, 1, 2)
	base[0] = Item{Name: "book"}
	view := base[:1]
	_ = AppendItem(view, Item{Name: "pen"})
	fmt.Printf("shared backing array side effect: base[1]=%s\n", base[:2][1].Name)

	full := make([]Item, 1, 1)
	full[0] = Item{Name: "notebook"}
	_ = AppendItem(full, Item{Name: "eraser"})
	fmt.Printf("no side effect when reallocated: len(full)=%d cap(full)=%d\n", len(full), cap(full))

	cl := ConsoleLogger{}
	UseLogger(cl) // implicit implementation

	processor := ProcessorImpl{logger: cl}
	badOrder := &Order{ID: "X", Total: 50}
	if err := SafeProcess(processor, badOrder); err != nil {
		fmt.Println("safe process error:", err)
	}
}
