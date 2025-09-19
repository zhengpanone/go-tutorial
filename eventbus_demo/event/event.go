package event

import (
	"sync"
)

type Event struct {
	EventType string
	Data      string
}

type EventBus struct {
	listeners []chan Event
	mu        sync.Mutex
}

func (eb *EventBus) AddListener(listener chan Event) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.listeners = append(eb.listeners, listener)
}

func (eb *EventBus) Dispatch(event Event) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	for _, listener := range eb.listeners {
		go func(l chan Event) {
			l <- event
		}(listener)
	}
}
