package persist

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"magro"
	"os"
	"testing"
	"time"
)

func TestPersister_Load(t *testing.T) {
	// given
	testPersistedData := Persisted{RecordedMacros: []*magro.Macro{{
		Name: "type name",
		Events: []magro.Event{
			{time.Microsecond, magro.KeyKindDown, 'a'},
			{time.Microsecond, magro.KeyKindUp, 'a'},
			{time.Microsecond, magro.KeyKindDown, 'l'},
			{time.Microsecond, magro.KeyKindUp, 'l'},
			{time.Microsecond, magro.KeyKindDown, 'e'},
			{time.Microsecond, magro.KeyKindUp, 'e'},
			{time.Microsecond, magro.KeyKindDown, 'x'},
			{time.Microsecond, magro.KeyKindUp, 'x'},
		},
	}}}

	file, err := os.CreateTemp("", "persisted.json")
	assert.NoError(t, err)

	fileContents, err := json.Marshal(&testPersistedData)
	assert.NoError(t, err)
	_, err = file.Write(fileContents)
	assert.NoError(t, err)

	persister, err := NewPersisterWithFile(file)
	assert.NoError(t, err)

	// when
	err = persister.Load()

	// then
	assert.NoError(t, err)
	assert.Equal(t, testPersistedData.RecordedMacros, persister.RecordedMacros)
}

func TestPersister_Start_updatesWhenMacrosChange(t *testing.T) {
	file, err := os.CreateTemp("", "persisted.json")
	assert.NoError(t, err)

	persister, err := NewPersisterWithFile(file)
	assert.NoError(t, err)

	changeCh := persister.Start()

	testMacro := &magro.Macro{Name: "cool testMacro"}
	persister.RecordedMacros = append(persister.RecordedMacros, testMacro)

	assert.Equal(t, Persisted{RecordedMacros: []*magro.Macro{testMacro}}, <-changeCh)
	assert.Zero(t, len(changeCh))
}
