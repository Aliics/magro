package magro

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"log"
	"time"
)

type Macro []hook.Event

func (m Macro) PlayEvents() error {
	log.Printf("playing %d m from %s", len(m), m[0].When)

	previousWhen := m[0].When
	for _, event := range m {
		time.Sleep(event.When.Sub(previousWhen))

		if event.Rawcode != 0 {
			fmt.Printf("kb event: %s\n", event)

			if event.Kind == hook.KeyDown {
				err := robotgo.KeyDown(string(rune(event.Rawcode)))
				if err != nil {
					return err
				}
			} else if event.Kind == hook.KeyUp {
				err := robotgo.KeyUp(string(rune(event.Rawcode)))
				if err != nil {
					return err
				}
			}
		}

		previousWhen = event.When
	}

	return nil
}
