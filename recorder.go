package magro

import (
	hook "github.com/robotn/gohook"
	"log"
	"time"
)

type Recorder struct {
	IsRecording    bool
	RecordedMacros *[]Macro

	recordCh    chan bool
	eventCh     chan hook.Event
	processedCh chan hook.Event

	previousEventTime *time.Time
	currentMacro      Macro
}

func NewRecorder(eventCh chan hook.Event) *Recorder {
	return &Recorder{
		RecordedMacros: &[]Macro{},

		recordCh:     make(chan bool),
		eventCh:      eventCh,
		processedCh:  make(chan hook.Event),
		currentMacro: NewMacro(),
	}
}

func (r *Recorder) Close() {
	close(r.recordCh)
}

func (r *Recorder) Toggle() {
	r.recordCh <- !r.IsRecording
}

func (r *Recorder) ToggleBlocking() {
	initialRecording := r.IsRecording
	r.Toggle()

	for r.IsRecording == initialRecording {
		// Loop until "IsRecording" has changed.
	}
}

func (r *Recorder) Start() <-chan hook.Event {
	go func() {
		for {
			select {
			case nowRecording := <-r.recordCh:
				if !nowRecording {
					*r.RecordedMacros = append(*r.RecordedMacros, r.currentMacro)
					r.currentMacro = NewMacro()
					r.previousEventTime = nil
				} else {
					log.Println("macro recording started")
					now := time.Now()
					r.previousEventTime = &now
				}

				// Reflect the recording state here.
				r.IsRecording = !r.IsRecording
			case event := <-r.eventCh:
				if r.IsRecording && event.Rawcode != 0 {
					r.addEventToMacroRecording(event)
				}

				r.processedCh <- event
			}
		}
	}()

	return r.processedCh
}

func (r *Recorder) addEventToMacroRecording(event hook.Event) {
	var delta time.Duration
	if r.previousEventTime != nil {
		delta = event.When.Sub(*r.previousEventTime)
	}

	var keyKind KeyKind
	if event.Kind == hook.KeyDown {
		keyKind = KeyKindDown
	} else {
		keyKind = KeyKindUp
	}

	macroEvent := Event{delta, keyKind, rune(event.Rawcode)}
	r.currentMacro.Events = append(r.currentMacro.Events, macroEvent)
	r.previousEventTime = &event.When

	log.Printf("recorded event %s", macroEvent)
}
