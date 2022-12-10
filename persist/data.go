package persist

import (
	"magro"
	"time"
)

type data struct {
	Macros []macro `json:"macros,omitempty"`
}

type macro struct {
	Name   string  `json:"name,omitempty"`
	Events []event `json:"events,omitempty"`
}

type event struct {
	DeltaMicros int64 `json:"deltaMicros,omitempty"`
	KeyKindNum  uint8 `json:"keyKindNum,omitempty"`
	Keycode     rune  `json:"keycode,omitempty"`
}

func persistedFromMacroList(macros []magro.Macro) data {
	var persistedMacros []macro
	for _, m := range macros {
		persistedMacros = append(persistedMacros, macro{
			Name:   m.Name,
			Events: persistedEventsFromEvents(m.Events),
		})
	}

	return data{persistedMacros}
}

func persistedEventsFromEvents(events []magro.Event) []event {
	var persistedEvents []event
	for _, e := range events {
		persistedEvents = append(persistedEvents, event{
			DeltaMicros: e.Delta.Microseconds(),
			KeyKindNum:  uint8(e.KeyKind),
			Keycode:     e.Keycode,
		})
	}

	return persistedEvents
}

func macroListFromPersisted(data data) []magro.Macro {
	var macros []magro.Macro
	for _, m := range data.Macros {
		macros = append(macros, magro.Macro{
			Name:   m.Name,
			Events: eventsFromPersistedEvents(m.Events),
		})
	}

	return macros
}

func eventsFromPersistedEvents(persistedEvents []event) []magro.Event {
	var events []magro.Event
	for _, e := range persistedEvents {
		events = append(events, magro.Event{
			Delta:   time.Duration(e.DeltaMicros) * time.Microsecond,
			KeyKind: magro.KeyKind(e.KeyKindNum),
			Keycode: e.Keycode,
		})
	}

	return events
}
