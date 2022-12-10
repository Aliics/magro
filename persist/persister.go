package persist

import (
	"encoding/json"
	"log"
	"magro"
	"os"
	"reflect"
)

const (
	configDirName   = ".magro"
	persistFileName = "persist.json"
)

type Persister struct {
	RecordedMacros *[]magro.Macro

	file *os.File

	isRunning bool
	changeCh  chan data
}

// NewPersister will create a new Persister with the internal file
// value being "~/.margo/persist.json". This file will be created if
// it does not already exist.
func NewPersister() (*Persister, error) {
	// Create or open the default file for persistence.
	// The full path will be created.
	var file *os.File
	{
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

		dirPath := homeDir + string(os.PathSeparator) + configDirName

		err = os.MkdirAll(dirPath, 0750)
		if err != nil && !os.IsExist(err) {
			return nil, err
		}

		filePath := dirPath + string(os.PathSeparator) + persistFileName

		file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return nil, err
		}
	}

	return NewPersisterWithFile(file)
}

func NewPersisterWithFile(file *os.File) (*Persister, error) {
	p := &Persister{
		RecordedMacros: &[]magro.Macro{},

		file:     file,
		changeCh: make(chan data),
	}

	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// Data exists, let's proceed early.
	if info.Size() > 1 {
		return p, err
	}

	// No data. Empty object.
	err = json.NewEncoder(p.file).Encode(persistedFromMacroList(*p.RecordedMacros))
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Persister) StartAndPersist() error {
	for range p.Start() {
		err := p.file.Truncate(0)
		if err != nil {
			return err
		}

		// Truncate doesn't reset offset.
		// Seek to the beginning.
		_, err = p.file.Seek(0, 0)
		if err != nil {
			return err
		}

		err = json.NewEncoder(p.file).Encode(persistedFromMacroList(*p.RecordedMacros))
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Persister) Start() <-chan data {
	p.isRunning = true

	go func() {
		var knownRecordedMacros []magro.Macro

		for p.isRunning {
			if reflect.DeepEqual(*p.RecordedMacros, knownRecordedMacros) {
				continue
			}

			log.Printf("persisting %d macros\n", len(*p.RecordedMacros))

			// Macros have changed.
			knownRecordedMacros = []magro.Macro{}
			knownRecordedMacros = append(knownRecordedMacros, *p.RecordedMacros...)

			p.changeCh <- persistedFromMacroList(*p.RecordedMacros)
		}
	}()

	return p.changeCh
}

func (p *Persister) Load() error {
	// Make sure that we reset the offset before attempting to load.
	_, err := p.file.Seek(0, 0)
	if err != nil {
		return err
	}

	var persisted data
	err = json.NewDecoder(p.file).Decode(&persisted)

	*p.RecordedMacros = append(*p.RecordedMacros, macroListFromPersisted(persisted)...)

	log.Printf("%d macros loaded", len(*p.RecordedMacros))

	return err
}

func (p *Persister) Close() {
	p.isRunning = false
	close(p.changeCh)
}
