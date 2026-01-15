package common

import (
	"fmt"
)

// NewBus creates and initializes a new event bus
func NewBus() *Bus {
	return &Bus{
		listeners: make(map[EventType][]Listener),
	}
}

// Subscribe registers a listener for a specific event type
func (b *Bus) Subscribe(t EventType, l Listener) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.listeners[t] = append(b.listeners[t], l)
}

// Publish sends an event to all listeners subscribed to its type
func (b *Bus) Publish(e GameEvent) {
	b.mu.RLock()
	listeners := b.listeners[e.Type]
	b.mu.RUnlock()

	fmt.Println("DEBUG: Publishing event:", e.Type, "to", len(listeners), "listeners")

	for _, l := range listeners {
		l(e)
	}
}

// Clear removes all listeners (useful for tests or resets)
func (b *Bus) Clear() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.listeners = make(map[EventType][]Listener)
}
