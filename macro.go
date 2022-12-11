package magro

import (
	"errors"
	"fmt"
	"github.com/go-vgo/robotgo"
	"log"
	"time"
)

var (
	errInvalidKeyKind = errors.New("unknown KeyKind")
)

type KeyKind uint8

func NewKeyKindFromString(s string) (KeyKind, error) {
	switch s {
	case "Down":
		return KeyKindDown, nil
	case "Up":
		return KeyKindUp, nil
	default:
		return 0, errInvalidKeyKind
	}
}

func (k KeyKind) String() string {
	switch k {
	case KeyKindDown:
		return "Down"
	case KeyKindUp:
		return "Up"
	default:
		panic(errInvalidKeyKind)
	}
}

const (
	KeyKindDown KeyKind = iota
	KeyKindUp
)

type Event struct {
	Delta   time.Duration
	KeyKind KeyKind
	Keycode rune
}

func (e Event) String() string {
	return fmt.Sprintf("Event{Delta: %s, KeyKind: %s, Keycode: %c}", e.Delta, e.KeyKind, e.Keycode)
}

type Macro struct {
	Name   string
	Events []Event
}

func NewMacro() Macro {
	return Macro{Name: "new macro"}
}

func (m Macro) PlayEvents() error {
	log.Printf("playing %d macro events", len(m.Events))

	for _, event := range m.Events {
		time.Sleep(event.Delta)

		if event.Keycode != 0 {
			log.Printf("kb event: %s\n", event)

			if event.KeyKind == KeyKindDown {
				err := robotgo.KeyDown(string(event.Keycode))
				if err != nil {
					return err
				}
			} else if event.KeyKind == KeyKindUp {
				err := robotgo.KeyUp(string(event.Keycode))
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
