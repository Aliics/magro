package magro

import (
	hook "github.com/robotn/gohook"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRecorder_simpleKeyStrokes(t *testing.T) {
	eventCh := make(chan hook.Event)
	recorder := NewRecorder(eventCh)
	runRecorder(recorder)

	awaitToggle(recorder)

	now := *recorder.previousEventTime
	recorder.eventCh <- keyDown(now, 'a')
	recorder.eventCh <- keyUp(now.Add(6*time.Microsecond), 'a')
	recorder.eventCh <- keyDown(now.Add(12*time.Microsecond), 's')
	recorder.eventCh <- keyUp(now.Add(14*time.Microsecond), 's')
	recorder.eventCh <- keyDown(now.Add(22*time.Microsecond), 'd')
	recorder.eventCh <- keyUp(now.Add(23*time.Microsecond), 'd')
	recorder.eventCh <- keyDown(now.Add(24*time.Microsecond), 'w')
	recorder.eventCh <- keyUp(now.Add(31*time.Microsecond), 'w')

	awaitToggle(recorder)

	assert.Equal(
		t,
		[]Macro{{
			Event{0 * time.Microsecond, KeyKindDown, 'a'},
			Event{6 * time.Microsecond, KeyKindUp, 'a'},
			Event{6 * time.Microsecond, KeyKindDown, 's'},
			Event{2 * time.Microsecond, KeyKindUp, 's'},
			Event{8 * time.Microsecond, KeyKindDown, 'd'},
			Event{1 * time.Microsecond, KeyKindUp, 'd'},
			Event{1 * time.Microsecond, KeyKindDown, 'w'},
			Event{7 * time.Microsecond, KeyKindUp, 'w'},
		}},
		recorder.RecordedMacros,
	)
}

func keyDown(when time.Time, r rune) hook.Event {
	return hook.Event{When: when, Rawcode: uint16(r), Kind: hook.KeyDown}
}

func keyUp(when time.Time, r rune) hook.Event {
	return hook.Event{When: when, Rawcode: uint16(r), Kind: hook.KeyUp}
}

func runRecorder(recorder *Recorder) {
	go func() {
		for range recorder.Start() {
			// Run without processing.
		}
	}()
}

func awaitToggle(recorder *Recorder) {
	initialRecording := recorder.IsRecording
	recorder.Toggle()

	for recorder.IsRecording == initialRecording {
		// Loop until "IsRecording" has changed.
	}
}
