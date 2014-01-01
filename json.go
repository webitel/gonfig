package gonfig

import (
	"encoding/json"
	"io/ioutil"
)

type JsonConfig struct {
	*MemoryConfig
	path string
}

// Returns a new Configurable backed by a json file at path.
// The file does not need to exist, if it does not exist the first Save call will create it.
func NewJsonConfig(path string) Configurable {
	cfg := &JsonConfig{&MemoryConfig{}, path}
	cfg.Load()
	return cfg
}

func (self *JsonConfig) Load() (err error) {
	if self.data == nil {
		self.initialize()
	}
	var data []byte = make([]byte, 1024)
	if data, err = ioutil.ReadFile(self.path); err != nil {
		return err
	}
	out, err := self.unmarshal(data)
	if err != nil {
		return err
	}
	self.data = out
	return nil
}

func (self *JsonConfig) Save() (err error) {
	if self.data == nil {
		self.initialize()
	}
	b, err := json.Marshal(self.data)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(self.path, b, 0600); err != nil {
		return err
	}

	return nil
}
