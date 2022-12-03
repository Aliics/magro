package magro

import (
	hook "github.com/robotn/gohook"
)

type Recorder struct {
	IsRecording    bool
	RecordedMacros [][]hook.Event

	recordCh    chan bool
	eventCh     chan hook.Event
	processedCh chan hook.Event

	currentEvents []hook.Event
}

func NewRecorder(eventCh chan hook.Event) *Recorder {
	return &Recorder{
		recordCh:    make(chan bool),
		eventCh:     eventCh,
		processedCh: make(chan hook.Event),
	}
}

func (r *Recorder) Close() {
	close(r.recordCh)
}

func (r *Recorder) Toggle() {
	r.recordCh <- !r.IsRecording
}

func (r *Recorder) Start() <-chan hook.Event {
	go func() {
		for {
			select {
			case nowRecording := <-r.recordCh:
				if !nowRecording {
					r.RecordedMacros = append(r.RecordedMacros, r.currentEvents)
					r.currentEvents = nil
				}

				// Reflect the recording state here.
				r.IsRecording = !r.IsRecording
			case event := <-r.eventCh:
				if r.IsRecording {
					r.currentEvents = append(r.currentEvents, event)
				}

				r.processedCh <- event
			}
		}
	}()

	return r.processedCh
}
