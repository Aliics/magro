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
	RecordedMacros []*magro.Macro

	file *os.File

	isRunning bool
	changeCh  chan Persisted
}

// NewPersister will create a new Persister with the internal file
// value being "~/.margo/persist.json". This file will be created if
// it does not already exist.
func NewPersister() (*Persister, error) {
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

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}

	return NewPersisterWithFile(file)
}

func NewPersisterWithFile(file *os.File) (*Persister, error) {
	p := &Persister{
		RecordedMacros: []*magro.Macro{},

		file:     file,
		changeCh: make(chan Persisted),
	}

	err := json.NewEncoder(p.file).Encode(Persisted{RecordedMacros: p.RecordedMacros})
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Persister) StartAndPersist() error {
	for persisted := range p.Start() {
		for _, macro := range persisted.RecordedMacros {
			p.RecordedMacros = append(p.RecordedMacros, macro)
		}

		err := json.NewEncoder(p.file).Encode(Persisted{RecordedMacros: p.RecordedMacros})
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Persister) Start() <-chan Persisted {
	p.isRunning = true

	go func() {
		var knownRecordedMacros []*magro.Macro

		for p.isRunning {
			nothingToPersist := len(p.RecordedMacros) == 0 && len(knownRecordedMacros) == 0
			if nothingToPersist || reflect.DeepEqual(p.RecordedMacros, knownRecordedMacros) {
				continue
			}

			log.Printf("persisting %d macros\n", len(p.RecordedMacros))

			// Macros have changed.
			for _, macro := range p.RecordedMacros {
				knownRecordedMacros = append(knownRecordedMacros, macro)
			}
			p.changeCh <- Persisted{RecordedMacros: p.RecordedMacros}
		}
	}()

	return p.changeCh
}

func (p *Persister) Load() error {
	_, err := p.file.Seek(0, 0)
	if err != nil {
		return err
	}

	var persisted Persisted
	err = json.NewDecoder(p.file).Decode(&persisted)

	for _, macro := range persisted.RecordedMacros {
		p.RecordedMacros = append(p.RecordedMacros, macro)
	}

	log.Printf("%d macros loaded", len(p.RecordedMacros))

	return err
}

func (p *Persister) Close() {
	p.isRunning = false
	close(p.changeCh)
}
