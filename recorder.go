package magro

import (
	hook "github.com/robotn/gohook"
)

type Recorder struct {
	isRecording bool

	recordCh    chan bool
	eventCh     chan hook.Event
	processedCh chan hook.Event

	recordedMacros [][]hook.Event
	currentEvents  []hook.Event
}

func NewRecorder(eventCh chan hook.Event) *Recorder {
	return &Recorder{
		recordCh:    make(chan bool),
		eventCh:     eventCh,
		processedCh: make(chan hook.Event),
	}
}

func (r *Recorder) Close() error {
	close(r.recordCh)
	return nil
}

func (r *Recorder) Toggle() {
	r.recordCh <- !r.isRecording
}

func (r *Recorder) Start() <-chan hook.Event {
	go func() {
		for {
			select {
			case nowRecording := <-r.recordCh:
				if !nowRecording {
					r.recordedMacros = append(r.recordedMacros, r.currentEvents)
					r.currentEvents = nil
				}

				// Reflect the recording state here.
				r.isRecording = !r.isRecording
			case event := <-r.eventCh:
				if r.isRecording {
					r.currentEvents = append(r.currentEvents, event)
				}

				r.processedCh <- event
			}
		}
	}()

	return r.processedCh
}
