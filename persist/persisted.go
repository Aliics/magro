package persist

import "magro"

type Persisted struct {
	RecordedMacros []*magro.Macro `json:"recordedMacros"`
}
