package fix44

import (
	"github.com/b2broker/simplefix-go/fix"
)

type EventsGrp struct {
	*fix.Group
}

func NewEventsGrp() *EventsGrp {
	return &EventsGrp{
		fix.NewGroup(FieldNoEvents,
			fix.NewKeyValue(FieldEventType, &fix.String{}),
			fix.NewKeyValue(FieldEventDate, &fix.String{}),
			fix.NewKeyValue(FieldEventPx, &fix.Float{}),
			fix.NewKeyValue(FieldEventText, &fix.String{}),
		),
	}
}

func (group *EventsGrp) AddEntry(entry *EventsEntry) *EventsGrp {
	group.Group.AddEntry(entry.Items())

	return group
}

type EventsEntry struct {
	*fix.Component
}

func makeEventsEntry() *EventsEntry {
	return &EventsEntry{fix.NewComponent(
		fix.NewKeyValue(FieldEventType, &fix.String{}),
		fix.NewKeyValue(FieldEventDate, &fix.String{}),
		fix.NewKeyValue(FieldEventPx, &fix.Float{}),
		fix.NewKeyValue(FieldEventText, &fix.String{}),
	)}
}

func NewEventsEntry() *EventsEntry {
	return makeEventsEntry()
}

func (eventsEntry *EventsEntry) EventType() string {
	kv := eventsEntry.Get(0)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (eventsEntry *EventsEntry) SetEventType(eventType string) *EventsEntry {
	kv := eventsEntry.Get(0).(*fix.KeyValue)
	_ = kv.Load().Set(eventType)
	return eventsEntry
}

func (eventsEntry *EventsEntry) EventDate() string {
	kv := eventsEntry.Get(1)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (eventsEntry *EventsEntry) SetEventDate(eventDate string) *EventsEntry {
	kv := eventsEntry.Get(1).(*fix.KeyValue)
	_ = kv.Load().Set(eventDate)
	return eventsEntry
}

func (eventsEntry *EventsEntry) EventPx() float64 {
	kv := eventsEntry.Get(2)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(float64)
}

func (eventsEntry *EventsEntry) SetEventPx(eventPx float64) *EventsEntry {
	kv := eventsEntry.Get(2).(*fix.KeyValue)
	_ = kv.Load().Set(eventPx)
	return eventsEntry
}

func (eventsEntry *EventsEntry) EventText() string {
	kv := eventsEntry.Get(3)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (eventsEntry *EventsEntry) SetEventText(eventText string) *EventsEntry {
	kv := eventsEntry.Get(3).(*fix.KeyValue)
	_ = kv.Load().Set(eventText)
	return eventsEntry
}
