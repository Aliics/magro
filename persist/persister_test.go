package persist

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPersister_Load(t *testing.T) {
	// given
	testPersistedData := data{Macros: []macro{{
		Name: "type name",
		Events: []event{
			{0, 0, 'a'},
			{0, 1, 'a'},
			{0, 0, 'l'},
			{0, 1, 'l'},
			{0, 0, 'e'},
			{0, 1, 'e'},
			{0, 0, 'x'},
			{0, 1, 'x'},
		},
	}}}

	file, err := os.CreateTemp("", "data.json")
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
	assert.Equal(t, testPersistedData, persistedFromMacroList(*persister.RecordedMacros))
}

func TestPersister_Start_updatesWhenMacrosChange(t *testing.T) {
	file, err := os.CreateTemp("", "data.json")
	assert.NoError(t, err)

	persister, err := NewPersisterWithFile(file)
	assert.NoError(t, err)

	changeCh := persister.Start()

	testData := data{[]macro{{Name: "cool testMacro"}}}
	*persister.RecordedMacros = append(*persister.RecordedMacros, macroListFromPersisted(testData)...)

	assert.Equal(t, testData, <-changeCh)
	assert.Zero(t, len(changeCh))
}
