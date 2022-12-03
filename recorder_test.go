package magro

import (
	hook "github.com/robotn/gohook"
	"github.com/stretchr/testify/assert"
	"github.com/vcaesar/keycode"
	"testing"
	"time"
)

func TestRecorder_simpleKeyStrokes(t *testing.T) {
	eventCh := make(chan hook.Event)
	recorder := NewRecorder(eventCh)
	runRecorder(recorder)

	awaitToggle(recorder)

	recorder.eventCh <- keyDown(time.UnixMilli(0), "a")
	recorder.eventCh <- keyUp(time.UnixMilli(5), "a")
	recorder.eventCh <- keyDown(time.UnixMilli(10), "s")
	recorder.eventCh <- keyUp(time.UnixMilli(15), "s")
	recorder.eventCh <- keyDown(time.UnixMilli(20), "d")
	recorder.eventCh <- keyUp(time.UnixMilli(25), "d")
	recorder.eventCh <- keyDown(time.UnixMilli(30), "w")
	recorder.eventCh <- keyUp(time.UnixMilli(35), "w")

	awaitToggle(recorder)

	assert.Equal(
		t,
		[][]hook.Event{{
			keyDown(time.UnixMilli(0), "a"),
			keyUp(time.UnixMilli(5), "a"),
			keyDown(time.UnixMilli(10), "s"),
			keyUp(time.UnixMilli(15), "s"),
			keyDown(time.UnixMilli(20), "d"),
			keyUp(time.UnixMilli(25), "d"),
			keyDown(time.UnixMilli(30), "w"),
			keyUp(time.UnixMilli(35), "w"),
		}},
		recorder.RecordedMacros,
	)
}

func keyDown(when time.Time, keyString string) hook.Event {
	return hook.Event{When: when, Keycode: keycode.Keycode[keyString], Direction: hook.KeyDown}
}

func keyUp(when time.Time, keyString string) hook.Event {
	return hook.Event{When: when, Keycode: keycode.Keycode[keyString], Direction: hook.KeyUp}
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
