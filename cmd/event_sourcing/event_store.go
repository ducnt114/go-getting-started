package event_sourcing

import "fmt"

type EventStore struct {
	events []Event
}

func NewEventStore() *EventStore {
	return &EventStore{events: []Event{}}
}

func (store *EventStore) Append(event Event) {
	store.events = append(store.events, event)
	fmt.Printf("Event stored: %s\n", event.EventType())
}

func (store *EventStore) GetAllEvents() []Event {
	return store.events
}
